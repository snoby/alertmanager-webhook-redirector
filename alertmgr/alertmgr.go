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

func (alert AlertMgr) AlertMgr_LoadRawData(buf []byte) (err error) {
	fmt.Println(" Loading buffer into AlertMgr memory")
	x := len(buf)
	fmt.Println(" Length of data : ", x)

	alert.buf = buf

	y := len(alert.buf)
	fmt.Println(" Length of alertMgr.buf : ", y)
	return nil
}

func (alert AlertMgr) AlertMgr_PrintRawJSON() (err error) {
	if alert.buf == nil {
		return errors.New("AlertMgr: No data loaded.  Load data first!!!")
	}

	var out bytes.Buffer
	err = json.Indent(&out, alert.buf, "", "  ")
	if err != nil {
		fmt.Print(err)
		return err
	}

	fmt.Printf("size of data: %v\n", (out.Bytes))

	fmt.Printf("%s", out.Bytes)
	return nil

}
