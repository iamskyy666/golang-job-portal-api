package utils

import (
	"fmt"
	"os"
)

func DeleteFileIfExists(path string) error {
	// Check if file exists
	if _,err:=os.Stat(path);os.IsNotExist(err){
		// File doesn't exist, return without error
		return nil
	}

	// Try deleting the file
	err:=os.Remove(path)
	if err!=nil{
		return fmt.Errorf("ERROR deleting the file: %v",err)
	}
	return nil
}


//func DeleteProfilePic(){}