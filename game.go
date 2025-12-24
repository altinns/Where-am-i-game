package main 
import(
	"fmt"
)



func main(){
	
player1:=[10]string{}
player2:=[10]string{}

var first string
var position1 int
var position2 int
var second string

fmt.Printf("player name:")
fmt.Scan(&first)
fmt.Printf("pick position from 0-10:")
fmt.Scan(&position1)


player1[position1]=first
for i:=0;i<10;i++{
	if player1[i]=="" {
		player1[i]="x"
	}
 
}
//fmt.Println(player1)
 
fmt.Printf("player name:")
fmt.Scan(&second)
fmt.Printf("pick position from 0-10:")
fmt.Scan(&position2)

player2[position2]=second



for i:=0;i<10;i++{
	if player2[i]=="" {
		player2[i]="x"
	}
 
}

fmt.Println(player1)
fmt.Println(player2)

p1Health:=8
//p2Health:=8

var p1Guess int
//var p2Guess string
fmt.Println("player1 guess:")
fmt.Scan(&p1Guess)
if player1[p1Guess]=="x"{
p1Health=p1Health-1
}
fmt.Println(p1Health)


}