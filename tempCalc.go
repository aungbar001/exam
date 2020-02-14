//Name:Aungelia Barltta
//date:2/12/20
//description: a calculator that converts a dagree into farinhight, celcius and kalvin
//cite: into to go language chapeter 3 

//pyodocode:
// so first i will give the equations for all the convertions that need to be done.
//then i will ask the user for a input for farinhight, calcius or kalven.
//then form there the calculator will take that number and run it threw the quation depending on what the user selected. 
package main

import "fmt"

func main(){
  //c2f=(c*5/9)+32
  //f2c=(f-32) * 5/9
  //c2k=(c+273.15)
  //k2f= (k - 273.15)*9/5+32
  //k2c=(k-273.15)
  //f2k=(f-32)*5/9+273.15
  
//asked user to give me a number 
  fmt.Print("Give a number")
  //named the varable number and had it float64 
  var number float64
  //
  fmt.Scanf("%f", &number)
  //these are asuming you know what the number you are entering is either farinhight, calcius, or kalvin degree 
  //witch then the output will spit out the right convertion for the one you are looking for
  output1 := (number * 9/5 + 32)

  output2 := (number - 32) * 5/9

  output3 := (number + 273.15)

  output4 := (number - 273.15) * 9/5 + 32

  output5 := (number - 273.15)

  output6 := (number -32) * 5/9 + 273.15
// now i named the convetion and print it on the screen so you know what your going too and form. and i have it spit out the numbe acrding to my calculations
  fmt.Println("Celsius to Farinhight:")
  fmt.Println(output1)

  fmt.Println("Farinhight to Celsius:")
  fmt.Println(output2)

  fmt.Println("Celsius to Kalvin:")
  fmt.Println(output3)

  fmt.Println("Kalvin to Farinhight:")
  fmt.Println(output4)

  fmt.Println("Kalvin to Celsius:")
  fmt.Println(output5)

  fmt.Println("Farinhight to Kalvin")
  fmt.Println(output6)








}
