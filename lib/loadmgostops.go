
package hsrgtfs

import (
   "fmt"
   "archive/zip"
   "encoding/csv"
   "strconv"
   "gopkg.in/mgo.v2"
)

func loadmgostops(f *zip.File){
   // f is our stop_times.txt
   fmt.Println("loadgmostops()");

   // open and read routes.txt
   rc, err := f.Open(); check(err); defer rc.Close();
   reader := csv.NewReader(rc);
   // read it line by line -records is an array
   records, err := reader.ReadAll(); check(err);


   
   // dial for a mongo db session
   sess, err := mgo.Dial("localhost"); check(err); defer sess.Close(); sess.SetMode(mgo.Monotonic, true);

   c := sess.DB("hsr").C("stops");

   // flush previous loading of routes
   c.DropCollection(); //ignore err cause we may not have a collection yet to drop if this is first run 
   


   for _, g := range records[1:len(records)]{

      // fmt.Println(g[1], g[4], g[6], g[2]);

      stoplat, err := strconv.ParseFloat(g[4], 64); check(err);
      stoplon, err := strconv.ParseFloat(g[5], 64); check(err);
      err = nil;
      stopcode, err := strconv.Atoi(g[1]); 
      stopid, err := strconv.Atoi(g[0]); 
      //fmt.Println(routeid, shapeid, tripid, directionid);
      if err == nil {
	 err = c.Insert( &Stop{
	    stopid,
	    stopcode,
	    g[2],
	    stoplat,
	    stoplon,
	 } ); check(err);
      }


   }
      
   // count of collection 
   //count, err := c.Count(); check(err);
   //fmt.Println("count: " , count);

}
