
package hsrgtfs

import (
   "fmt"
   "archive/zip"
   "encoding/csv"
   "strconv"
   "gopkg.in/mgo.v2"
)

func loadmgostoptimes(f *zip.File){
   // f is our stop_times.txt
   fmt.Println("loadgmostoptimes()");

   // open and read routes.txt
   rc, err := f.Open(); check(err); defer rc.Close();
   reader := csv.NewReader(rc);
   // read it line by line -records is an array
   records, err := reader.ReadAll(); check(err);


   
   // dial for a mongo db session
   sess, err := mgo.Dial("localhost"); check(err); defer sess.Close(); sess.SetMode(mgo.Monotonic, true);

   c := sess.DB("hsr").C("stoptimes");

   // flush previous loading of routes
   c.DropCollection(); //ignore err cause we may not have a collection yet to drop if this is first run 
   


   for _, g := range records[1:len(records)]{

      // fmt.Println(g[1], g[4], g[6], g[2]);

      tripid, err := strconv.Atoi(g[0]); check(err);
      err = nil;
      stopid, err := strconv.Atoi(g[3]); // skip on error 
      
     // stopid gives errors so lets skip it
     // sometimes instead of an int it is
     // something like 1626_merged_928740
     // so we'll just skip those for now
     if err == nil{
	fmt.Println(tripid, stopid, g[1], g[2]);

	err = c.Insert( &Stoptime{
	   tripid,
	   g[1], // arrival_time
	   g[2], // departure_time
	   stopid,
	} ); check(err);

     }

   }

      
   // count of collection 
   // fmt.Println("calculate count");
   // count, err := c.Count(); check(err);
   // fmt.Println("count: " , count);

}
