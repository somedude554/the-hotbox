# the-hotbox
Documentation of my progress on learning HTML.
I learned html at https://www.codecademy.com/learn/learn-html
Important notes:
1. Opening and closing tags. HTML has tags that help structure and format the coding of the website to make it easier to read and understand. Displayed content is placed within a tag. An example of an opening tag is `<p>` and a closing tag is `</p>`.
```
<p>Hello World</p>
```
2. The body tags, `<body>` and `</body>` are used to construct the body of a website. You are able to place other tags within the body tag like so:
```
<body>
  <p>"Life is very short and what we have to do must be done in the now." - Audre Lorde</p>
</body>
```
3. The structure is important for easy readibilty. Opening tags and closing tags have the same indentions and nested tabs are indented further than the parent tabs like the example above.
4. There exists heading tabs that help structure the website: h1-h6. h1 appears larger than h2 and so on.
```
<body>
  <h1>The Brown Bear</h1>
  <h2>About Brown Bears</h2>
  <h3>Species</h3>
  <h3>Features</h3>
  <h2>Habitat</h2>
  <h3>Coutries with Large Brown Bear Populations</h3>
  <h3>Countries with Small Brown Bear Populations</h3>
  <h2>Media</h2>
</body>
```
5. The div tag helps group certain elements together that one wants grouped, making it easier to modify a section of code rather than indivdual lines.
```
<body>
  <h1>The Brown Bear</h1>
  <div>
    <h2>About Brown Bears</h2>
    <h3>Species</h3>
    <h3>Features</h3>
  </div>
  <div>
    <h2>Habitat</h2>
    <h3>Countries with Large Brown Bear Populations</h3>
    <h3>Countries with Small Brown Bear Populations</h3>
  </div>
  <div>
    <h2>Media</h2>
  </div>
</body>
```
6. Attributes are used to expand an elements tag. Attributes are made up of a name and a value. For example, "id" can be used to specify different content within the code.
```
<body>
  <h1>The Brown Bear</h1>
  <div id = "introduction">
    <h2>About Brown Bears</h2>
    <h3>Species</h3>
    <h3>Features</h3>
  </div>
  <div id = "habitat">
    <h2>Habitat</h2>
    <h3>Countries with Large Brown Bear Populations</h3>
    <h3>Countries with Small Brown Bear Populations</h3>
  </div>
  <div id = "media">
    <h2>Media</h2>
  </div>
</body>
```
7. The `<p>` tag is for paragraph and allows for text to be formatted and separated based on if it's in the paragraph or not.
```
 <div>
  <h1>Technology</h1>
</div>
<div>
  <p><span>Self-driving cars</span> are anticipated to replace up to 2 million jobs over the next two decades.</p>
</div>
```
8. You can style text using the `<em>` for italics and `<strong>` for bold.
9. Use `<br>` for line breaks.
10. Use `<ul>` and `<ol>` for unordered and ordered list respectively. THe listed items must use the `<li>` tag.
```
<ul>
  <li>Limes</li>
  <li>Tortillas</li>
  <li>Chicken</li>
</ul>
```
11. Add images with the `<img>` tag. Use the keyword "src" to designate the location of the image. The alt attribute is also included to describe the image. It does not require a closing tag.
```
<img src="image-location.jpg" />
<img src="#" alt="A field of yellow sunflowers" />
```
12. Videos can be added using the <video> tag and also requires an "src" attribute. Video does require a closing tag like most other tags. The text in between the tags can display an error message if the video can not properly load.
```
<video src="myVideo.mp4" width="320" height="240" controls>
  Video not supported
</video>
```
13. Many of the html elements can call script that executes code that can also modify any of the CSS or html elements. Example from http://www.simplehtmlguide.com/javascript.php
```
<html>
 <head>
  <script type="text/javascript">
function functionOne() { alert('You clicked the top text'); }
function functionTwo() { alert('You clicked the bottom text'); }
  </script>
 </head>
<body>
 <p><a href="#" onClick="functionOne();">Top Text</a></p>
 <p><a href="javascript:functionTwo();">Bottom Text</a></p>
 </body>
</html>
```
  
I also learned some CSS at https://developer.mozilla.org/en-US/docs/Web/CSS/Using_CSS_custom_properties
Important Notes:
1. Classes are created with the ".". Example: ".light-mode" denotes the light-mode class.
2. There are many options to customize the class: color, background color, margin, width, height, and display. Usage:
```
.someclass {
  color: white;
  background-color: brown;
  margin: 10px;
  width: 50px;
  height: 50px;
  display: inline-block;
}
```
3. Custom variables are allowed and are denoted by "--".
```
.someclass{
  --secondary-color: 
}
```
4. Customization of attribute for every <p> element of a parent. For the example below, it takes every odd element of row and changes the color to the specified second color.
https://www.w3schools.com/cssref/sel_nth-child.asp
```
.row:nth-child(odd){
  color:var(--secondary-color);
}
```

I refreshed my JavaScript knowledge at https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/JavaScript_basics
JavaScript is incredibly useful in frontend development, since it allows for a user to dynamically change both the html or css directly from the code.
Since the syntax and usage of JavaScript is similar to Java and other object-oriented programming languages, most of the knowledge was easily applicable since I have experience with Java, C++, and python.
1. Initialize variables with the "var" keyword. These can also be modified later. Also end all lines with a semicolon.
```
var myVariable;
myVariable = 5;
```
2. The data types in JS are similar to Java: strings, numbers, booleans, arrays, and objects.
3. Multi-line comments are denoted using ``` /* ``` and ```*/``` and single line comments are denoted using ```//```.
4. Operators in JS are essentially the same as other languages: addition, subtraction, multiplication, division, assignment, equality, not, does not equal.
5. Functions are denoted with the "function" keyword. Parameters are specified by the programmer.
```
function multiply(num1,num2){
  return num1*num2;
}
```
6. JS functions can dynamically change the elements' attributes of the html and the css. Refer the the example below which is a snippet of the toggleTheme function in the main repository that I wrote.
```
function toggleTheme(){
  var element = document.body;
  var themeID = document.getElementById("themeSwitch").value;
  if(themeID == 1){
    document.documentElement.className="light-mode";
    document.getElementsByTagName("h1")[0].style.color = "black";
    document.getElementsByTagName("footer")[0].style.color = "black";
  }
}
```
The function ``` toggleTheme() ``` is being defined.
``` var element = document.body;``` sets the local variable "element" as the document body.
``` var themeID = document.getElementById("themeSwitch").value;``` sets the local variable "themeID" as the value of the "themeSwitch" element within the HTML.
The if statement checks if themeID is equivalent to 1 and if it is, it changes the CSS class to "light-mode", the 0th element of header to black, and the 0th element of footer to black.
7. JS functions can also take in outside data and modify the HTML/CSS accordingly. Snippet of code taken from the main HotBox repository.
References from https://www.w3schools.com/js/js_date_methods.asp
```
function checkDate(){
  var date = new Date();
  var month = date.getMonth()+1;
  if(month == 10){
    document.getElementById("themeSwitch").value=3;
  }
  toggleTheme();
}
```
The function ``` checkDate() ``` is being defined. This function checks the date and sets the theme according to the month.
``` var date = new Date();``` creates a new Date object as the "date" variable. ``` new Date()``` is a built in function in JS that takes the current date.
``` var month = date.getMonth()+1;``` sets the month as the current numberical month, e.g. October being 10. The ```.getMonth()``` method takes the current month and adds by 1 since January starts at 0. 
The if statement checks if the month is 10, or October, and if it is it changes the "themeSwitch" HTML element to 3, which refers to the Halloween Theme, and then terminates the loop. "toggleTheme()" function is then called that dynamically changes the theme of the website.
