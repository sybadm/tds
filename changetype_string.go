// Code generated by "stringer -type=changeType"; DO NOT EDIT.

package tds

import "strconv"

const _changeType_name = "dbChangelangChangecharsetChangepacketSizeChange"

var _changeType_index = [...]uint8{0, 8, 18, 31, 47}

func (i changeType) String() string {
	i -= 1
	if i >= changeType(len(_changeType_index)-1) {
		return "changeType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _changeType_name[_changeType_index[i]:_changeType_index[i+1]]
}