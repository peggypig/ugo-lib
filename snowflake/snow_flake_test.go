package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	sfc := SnowFlakeConfig{
		//2018-12-29 00:00:00
		StartTime:        1546012800000,
		DataCenterId:     0,
		WorkerId:         0,
		DataCenterIdBits: 8,
		WorkerIdBits:     6,
		SequenceBits:     8,
	}
	sf := SnowFlake{}
	err := sf.SetConfig(sfc)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 1000; i++ {
		fmt.Println(sf.Snow())
	}
}
