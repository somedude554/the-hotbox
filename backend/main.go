package main

import (
	"fmt"
	"strings"
	"os"
	"gocv.io/x/gocv"
)

// method;
// adds/subtracts a constant value from each pixel, modifying the brightness
func ModifyBrightness(frame *gocv.Mat, change uint8, inc bool) {
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

// this method does not work, but serves as a test for getting and setting using 
// GetUCharAt3 and SetUCharAt3; these methods don't traverse channels and therefore 
// this isn't a useful method. I am leaving this in and commenting it for learning purposes
func ModifyBrightness2(frame *gocv.Mat, change uint8, inc bool) {
	if change == 0 { return }
	
	for i := 0; i < frame.Rows(); i++ {
		for j := 0; j < frame.Cols(); j++ {
			
			for channel := 0; channel < 3; channel++ {
				
				// newval will eventually be the replacement value
				var newval uint8 
				var temp int16
				
				// store the result of addition/subtraction into temp
				if inc { 
					temp = int16(frame.GetUCharAt3(i, j, channel)) + int16(change) 
				} else { 
					temp = int16(frame.GetUCharAt3(i, j, channel)) - int16(change) 
				}
	
				// cap the brightness of each pixel between 0 and 255
				if temp > 255 {
					newval = 255	
				} else if temp < 0 {
					newval = 0
				} else {
					newval = uint8(temp)
				}
				
				// set the channel of the pixel to the calculated value
				frame.SetUCharAt3(i, j, channel, newval)

				// NOTE this doesn't work because you can't select channels using SetUCharAt3
			}

		}
	}
}

// this is what I prefer the most, using the data pointer itself to quickly edit the data
func ModifyBrightness3(frame *gocv.Mat, change int16) {
	framedata := frame.DataPtrUint8()

	// nice closure to expedite the process of keeping the values between 0 and 255
	limit := func(val int16) uint8 { 
		if val < 0 {
			return 0
		} else if val > 255 {
			return 255
		} else {
			return uint8(val)
		}
	}

	// framedata is one long array of uint8's, every third item represents a new pixel
	// and the three in between are the BGR channels
	for i := 0; i < len(framedata); i += 3 {
		// done like this so we can add weights
		framedata[i] = limit(int16(framedata[i]) + change) // B 
		framedata[i+1] = limit(int16(framedata[i+1]) + change) // G
		framedata[i+2] = limit(int16(framedata[i+2]) + change) // R
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("arg err") // check number of cli arguments
		return
	}

	file := os.Args[1] // get file name

	fileinfo, err := os.Stat(file) 
	if err != nil {
		fmt.Println("file err") // check file exits
		return
	}
	size := fileinfo.Size()

	if size > 250000000 {
		fmt.Println("file too large") // check file size
		return
	}

	video, _ := gocv.VideoCaptureFile(file) // open file as video
	defer video.Close()

	// Read the first frame, pass it to the JS API
	first := gocv.NewMat() // reader mat
	defer first.Close()
  	vidcap := gocv.VideoCapture(*video)     
  	success := vidcap.Read(&first)     
  	if (success) {
  		gocv.IMWrite("first_frame.jpg", first)  // save frame as JPEG file
  	}

	var outfilename strings.Builder
	outfilename.WriteString(file[:strings.Index(file,".")])
	outfilename.WriteString("out")
	outfilename.WriteString(file[strings.Index(file,"."):])

	out, _ := gocv.VideoWriterFile(outfilename.String(),
								video.CodecString(),
								video.Get(gocv.VideoCaptureFPS),
								int(video.Get(gocv.VideoCaptureFrameWidth)),
								int(video.Get(gocv.VideoCaptureFrameHeight)),
								true)
	defer out.Close()

	curr := gocv.NewMat() // reader mat
	defer curr.Close()
	
	for {
		// the frame curr is of type CV8UC3
		// CV8U means it stores unsigned chars (lit. 8 bit Unsigned)
		// C3 means it has 3 channels, these channels represent the Blue Green and Red respectively
		if ok := video.Read(&curr); !ok {
			fmt.Println("video reading stopped") // read frame to reader mat
			
			return
		}
		
		if curr.Empty() { 
			continue
		}
		
		// function call;
		// takes in a Mat*, a change value, and a boolean that asks if you're increasing or decreasing
		// changes the frame pointed to by the Mat*
		//ModifyBrightness(&curr, 50, false) // this reduces curr's brightness by 50
		ModifyBrightness3(&curr, -50)
		out.Write(curr)

	}

	// NOTE: the output doesn't have sound
}
