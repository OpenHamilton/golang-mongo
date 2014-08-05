// nadeem web design
// nad@webscripts.biz 
// consuming hamilton bus system gtfs data

package main

import (
   "nad/lib"
)



func main() {

   // download the zip file
   hsrgtfs.Downloadzipfile();

   // extract the data and load into mongo db
   hsrgtfs.Loadmgo();

}
