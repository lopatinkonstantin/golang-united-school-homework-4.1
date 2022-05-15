package reverse_string

func ReverseString(input string) (output string) {
	output=""
   	r:=[]rune(input)
   	for i:=len(r)-1;i>=0;i--{
   		output=output+string(r[i])
   	}
   	return output
}
