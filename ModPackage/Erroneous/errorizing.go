package erroneous;  

import (
   "errors"
)

func ThrowCustomError(message string) error {
	return errors.New(message); 
} 


