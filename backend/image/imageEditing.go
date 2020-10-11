package image

import (
	"math"
	"gocv.io/x/gocv"
	"../utils"
)


func ModifyBrightness(frame *gocv.Mat, change int16) {
	framedata := frame.DataPtrUint8()
	// framedata is one long array of uint8's, every third item represents a new pixel
	// and the three in between are the BGR channels
	for i := 0; i < len(framedata); i += 3 {
		// done like this so we can add weights
		framedata[i] = utils.Int16ToUint8(int16(framedata[i]) + change) // B 
		framedata[i+1] = utils.Int16ToUint8(int16(framedata[i+1]) + change) // G
		framedata[i+2] = utils.Int16ToUint8(int16(framedata[i+2]) + change) // R
	}
}


const MAXIMUM_BRIGHTNESS = 3
func ModifyContrast(frame *gocv.Mat, alpha float64) {
	framedata := frame.DataPtrUint8()
	
	// precomputes all brightness for this value for alpha
	var precomputed_brightness [256]float64
	for i := 0; i < 256; i++ {
		precomputed_brightness[i] = 255*(1 - 1/(1 + math.Pow(255.0/float64(i) - 1, -MAXIMUM_BRIGHTNESS*alpha - 1)))
	}

	// goes through every pixel and does the following:
	/*		calculates the highest brightness of any color channel in the pixel
	 *		
	 * 		finds the value that the brightness maps to, find out by how much it's scaled
	 * 		scales every channel accordingly
	 */
	for i := 0; i < len(framedata); i += 3 {
		value := utils.Max(framedata[i], framedata[i+1], framedata[i+2])

		var factor float64 = precomputed_brightness[value] / float64(value)
		for j := 0; j < 3; j++ {
			framedata[i+j] = uint8(factor * float64(framedata[i+j]))
		}
	}
}



// Less computation and faster than ModifySaturation1. Explained in the link below
// https://stackoverflow.com/questions/13806483/increase-or-decrease-color-saturation
// -1 <= scale <= 1
func ModifySaturation(frame *gocv.Mat, scale float64) {
	framedata := frame.DataPtrUint8()
	for i := 0; i < len(framedata); i += 3 {
		b := float64(framedata[i])
		g := float64(framedata[i+1])
		r := float64(framedata[i+2])
		gray := 0.1140*b + 0.5870*g + 0.2989*r

		framedata[i] = utils.Float64ToUint8(-gray*scale + b*(1+scale))
		framedata[i+1] = utils.Float64ToUint8(-gray*scale + g*(1+scale))
		framedata[i+2] = utils.Float64ToUint8(-gray*scale + r*(1+scale))
	}
}





// The most standard way to modify saturation, but it's slow because of color space conversion
// -1 <= scale <= 1
func ModifySaturation1(frame *gocv.Mat, scale float64) {
	//Convert a BGR image to a HSV image
	gocv.CvtColor(*frame, frame, gocv.ColorBGRToHSV)

	framedata := frame.DataPtrUint8()

	for i := 0; i < len(framedata); i += 3 {
		//Modify S(saturation) only
		new_saturation := float64(framedata[i+1]) * (1+scale)
		if new_saturation > 255 {
			new_saturation = 255
		}

		framedata[i+1] = uint8(new_saturation)
	}

	gocv.CvtColor(*frame, frame, gocv.ColorHSVToBGR)

}


// adds/subtracts a constant value from each pixel, modifying the brightness
func ModifyBrightness1(frame *gocv.Mat, change uint8, inc bool) {
	if change == 0 { return }

	// channels is of type []Mat, each of the three channels of frame
	// are now their own Mat type, and we can work with them separately
	channels := gocv.Split(*frame)

	// now we go through each channel and add the value of change to each pixel
	for i := 0; i < 3; i++ {
		// POSSIBLE OPTIMIZATION
		// Because we'll be doing multiple video operations at a time,
		// going through the frame as many times as we have adjustments to make is inefficient
		// if we could do this pixel-by-pixel we could do all of the adjustments at the same time
		// for now this is good but if we find our code to be too slow then we can come back to this

		if inc {
			channels[i].AddUChar(change) // channel i += change
		} else {
			channels[i].SubtractUChar(change) // channel i -= change
		}
	}

	// merges all the Mat's in channels into one multi-channel Mat, that being frame
	// now that we have added the value to the channels, we can put them back together in the frame
	gocv.Merge(channels, frame)

	// no return, the changes happen to the frame's reference
}