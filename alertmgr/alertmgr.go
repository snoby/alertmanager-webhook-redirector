package alertmgr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type AlertMgr struct {
	buf []byte
}

func (alert *AlertMgr) AlertMgr_LoadRawData(buf []byte) (err error) {
	fmt.Println(" Loading buffer into AlertMgr memory")
	alert.buf = buf
	return nil
}

func (alert *AlertMgr) AlertMgr_PrintRawJSON() (err error) {
	if alert.buf == nil {
		return errors.New("AlertMgr: No data loaded.  Load data first!!!")
	}

	var out bytes.Buffer
	err = json.Indent(&out, alert.buf, "", "  ")
	if err != nil {
		fmt.Println("Not able to Indent the buffer properly")
		fmt.Print(err)
		return err
	}

	fmt.Printf("size of data: %d\n", out.Len())
	fmt.Printf("%s\n", out.String())

	return nil

}
