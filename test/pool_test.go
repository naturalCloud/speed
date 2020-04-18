package test_test

import (
	"fmt"
	"speed/app/lib/faker"
	"speed/app/lib/workerPool"
	app "speed/bootstrap"
	"testing"
	"time"
)

func TestPool(t *testing.T) {

	p := workerPool.NewPool(5)
	ts := &workerPool.Task{}
	ts.SetHandel(func() error {
		fmt.Println("打酱油任务..........")
		return nil
	})

	//p.ExternalChain <- ts

	go func() {
		for i := 0; i < 50; i++ {
			p.ExternalChain <- ts
		}
	}()

	go p.Run()
	select {}

}

func BenchmarkInserData(b *testing.B) {
	facke := faker.NewFaker("/home/zhangshuai/project/go/speed/resources/data/faker/")
	type Data struct {
		ID        uint `gorm:"primary_key"`
		Name      string
		Card_id   string
		Phone     string
		Bank_id   string
		Email     string
		Address   string
		Create_at int
	}

	model := app.Db.Table("data")

	pool := workerPool.NewPool(50)

	start := time.Now().Unix()
	fmt.Println(start)

	go pool.Run()

	i := 0
	for ; i < b.N; i++ {

		task := workerPool.Task{}
		task.SetHandel(func() error {
			cid, _ := facke.MakeIdentificationCard()
			name, _ := facke.MakeName()
			dData := Data{
				Name:      name,
				Card_id:   cid,
				Phone:     facke.MakeMobile(),
				Bank_id:   facke.MakeBankCardId(),
				Email:     facke.MakeEmail(),
				Address:   facke.MakeAddress(),
				Create_at: int(time.Now().Unix()),
			}
			model.Create(&dData)
			return nil
		})
		pool.ExternalChain <- &task

	}

	fmt.Println("投递完成-------", time.Now().Unix()-start, i)




}
