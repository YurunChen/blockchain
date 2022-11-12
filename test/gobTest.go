package main

type Person struct {
	Name string
	Age  uint
}

//func main() {
//	XiaoMing := &Person{
//		"xiaoming",
//		20,
//	}
//	var buffer bytes.Buffer
//	encoder := gob.NewEncoder(&buffer)
//	err := encoder.Encode(XiaoMing)
//	if err != nil {
//		log.Panic("编码错误！")
//	}
//	fmt.Printf("编码后的小明：%v\n", buffer.Bytes())
//	var DaMing Person
//	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
//	err = decoder.Decode(&DaMing)
//	if err != nil {
//		log.Panic("解码错误！")
//	}
//	fmt.Printf("编码后的大明：%v\n", DaMing)
//}
