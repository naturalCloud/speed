package test

import (
	"fmt"
	"github.com/syyongx/php2go"
	"speed/app/lib/faker"
	"testing"
)

func TestBankId(t *testing.T) {
	newFaker := faker.NewFaker("/home/zhangshuai/project/go/speed/resources/data/faker/")
	name, err2 := newFaker.MakeName()
	if err2 != nil {
		t.Fatal(err2)
	}

	email := newFaker.MakeEmail()
	id := newFaker.MakeBankCardId()
	card, err := newFaker.MakeIdentificationCard()

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(id)
	fmt.Println(card)
	fmt.Println(newFaker.MakeMobile())
}

func BenchmarkBankIdB(b *testing.B) {
	newFaker := faker.NewFaker("/home/zhangshuai/project/go/speed/resources/data/faker/")

	for i := 0; i < b.N; i++ {
		 newFaker.MakeName()
		newFaker.MakeEmail()
		newFaker.MakeBankCardId()
		newFaker.MakeIdentificationCard()

	}

}

func BenchmarkTestc(b *testing.B) {
	var mobileSegment = []string{
		"133", "153", "180", "181", "189", "177", "173", "149",
		"130", "131", "132", "155", "156", "145", "185", "186", "176",
		"175", "135", "136", "137", "138", "139", "150", "151", "152",
		"157", "158", "159", "182", "183", "184", "187", "147", "178",
	}

	i := php2go.Rand(0, len(mobileSegment)-1)
	if i <= 0 {
		b.Fatal(i)
	}
	fmt.Println(i)
	//mobile := mobileSegment[php2go.Rand(0, len(mobileSegment))-1]
}

func TestAddress(t *testing.T)  {
	newFaker := faker.NewFaker("/home/zhangshuai/project/go/speed/resources/data/faker/")
	address := newFaker.MakeAddress()
	fmt.Println(address)
}

func BenchmarkAddress(b *testing.B) {
	newFaker := faker.NewFaker("/home/zhangshuai/project/go/speed/resources/data/faker/")

	for i := 0; i < b.N ; i++ {
		address := newFaker.MakeAddress()
		if address == "" {
			b.Error("错误")
		}
	}
}