package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"os"
)

// 业务代码
func yewuTask1(sqlStr string) error {
	if err := basetask1(sqlStr); err != nil {
		msg := fmt.Sprintf("sql: %s, func: %s", sqlStr, "task1")
		return errors.Wrap(err, msg)
	}
	return nil
}

// 业务代码
func yewuTask2(sqlStr string) error {
	if err := yewuTask1(sqlStr); err != nil {
		msg := fmt.Sprintf("sql: %s, func: %s", sqlStr, "task1")
		return errors.WithMessage(err, msg)
	}
	return nil
}

// 基础组件代码
func basetask1(sqlStr string) (err error) {
	// do base task
	// do ... error
	log.Println(sqlStr)
	// raise err
	score := rand.Intn(100)
	if score >= 100 {
		err = sql.ErrNoRows
	} else {
		err = os.ErrClosed
	}
	return
}

func main() {
	sqlStr := "select * from t1 where 1=2"
	err := yewuTask2(sqlStr)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			fmt.Printf("data not found,  %+v\n", err)
			return
		} else {
			log.Fatalln(err.Error())
			return
		}
		// unknown error
	}

}
