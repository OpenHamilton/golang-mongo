
package hsrgtfs 

import (
      "fmt"
      "net/http"
      "io/ioutil"
      )

func Downloadzipfile(){
   fmt.Println("Downloadzipfile()");

   resp, err := http.Get(Zipwebaddress);
   check(err); 
   defer resp.Body.Close();

   body, err := ioutil.ReadAll(resp.Body); 
   check(err);
   fmt.Println("hsr gtfs feed download complete");

   ioutil.WriteFile(Zippath, body, 0644);
   fmt.Println("hsr gtfs feed written to file hsrgtfsfeed.zip");
   fmt.Println("path ", Zippath);

}

