
package hsrgtfs

import (
   "fmt"
   "archive/zip"
   "encoding/csv"
   "strconv"
   "gopkg.in/mgo.v2"
)

func loadmgotrips(f *zip.File){
   // f is our trips.txt
   fmt.Println("loadgmotrips()");

   // open and read routes.txt
   rc, err := f.Open(); check(err); defer rc.Close();
   reader := csv.NewReader(rc);
   // read it line by line -records is an array
   records, err := reader.ReadAll(); check(err);


   
   // dial for a mongo db session
   sess, err := mgo.Dial("localhost"); check(err); defer sess.Close(); sess.SetMode(mgo.Monotonic, true);

   c := sess.DB("hsr").C("trips");

   // flush previous loading of routes
   c.DropCollection(); //ignore err cause we may not have a collection yet to drop if this is first run 
   


   for _, g := range records[1:len(records)]{

      // fmt.Println(g[1], g[4], g[6], g[2]);

      routeid, err := strconv.Atoi(g[1]); check(err);
      shapeid, err := strconv.Atoi(g[4]); check(err);
      tripid, err := strconv.Atoi(g[6]); check(err);
      directionid, err := strconv.ParseBool(g[2]); check(err);
      //fmt.Println(routeid, shapeid, tripid, directionid);
      
      err = c.Insert( &Trip{
	 routeid,
	 shapeid,
	 tripid,
	 directionid,
      } ); check(err);
      


   }
      
   // count of collection 
   // count, err := c.Count(); check(err);
   // fmt.Println("count: " , count);

}
