package main



func generica(interf interface{}){

}

func main(){
generica("texto string")
generica(32)
generica(true)
generica(14.4)
}