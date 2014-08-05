
package hsrgtfs

import (
   "fmt"
   "archive/zip"
   "encoding/csv"
   "strconv"
   "sort"
   "gopkg.in/mgo.v2"
)

func loadmgoroutes(f *zip.File){
   // f is our routes.txt
   fmt.Println("loadgmoroutes()");

   // open and read routes.txt
   rc, err := f.Open(); check(err); defer rc.Close();
   reader := csv.NewReader(rc);
   // read it line by line -records is an array
   records, err := reader.ReadAll(); check(err);

   // create an array to hold our route_id's for ordering
   ascendrouteids := make([]int, len(records)-1);

   // create a map to hold a map by route_id's 
   byrouteid := make(map[int][]string);

   // iterate over each line of routes.txt
   for i, g := range records[1:len(records)]{
      // convert our route_id string to an int so we can easily use it for ordering
      routeid, _ := strconv.Atoi(g[5]);
      
      // add each route array to our map
      // g[8] is short name ex. 08,
      // g[0] is long name ex. York
      // g[5] is id ex. 2903
      byrouteid[routeid] =  []string{ g[8], g[0], g[5] } ;

      // fill array to be used for ordering
      ascendrouteids[i] = routeid;

   }
   // order the list of routes by route_id
   sort.Ints(ascendrouteids);


   // print out our list for inspection
   /*
   for _, h := range ascendrouteids {
      for _, g := range records {
	 if id, _ := strconv.Atoi(g[5]); id == h {
	    fmt.Println(h, g); 
	 }
      }
   }
   */
   // routes are ordered from outbound to inbound 
   // starting with 01 KING 2792 assume direction 0/false
   // then repeating starting with 01 KING 2834 direction 1/true
   // so lets mark them by DirectionId 
   // but not sure exactly what it refers to 
   // E/W, N/S, Outbound/Inbound, corresponds to Trips DirectionId?


   // find index of 01 KING 2834 for splitting the list
   var index2834 int;
   for i, g := range ascendrouteids {
      if g == 2834 {
	 index2834 = i;
      }
   }


   // dial for a mongo db session
   sess, err := mgo.Dial("localhost"); check(err); defer sess.Close(); sess.SetMode(mgo.Monotonic, true);

   c := sess.DB("hsr").C("routes");

   // flush previous loading of routes
   c.DropCollection(); //ignore err cause we may not have a collection yet to drop if this is first run 

   err = c.Insert( &SplitRoute{
      ascendrouteids[:index2834], // Outbound 0
      ascendrouteids[index2834:], // Inbound 1
   } ); check(err);
      

   for _, rtid := range ascendrouteids[:index2834] {

      intid, err := strconv.Atoi(byrouteid[rtid][2]); check(err)
      intnumber, err := strconv.Atoi(byrouteid[rtid][0]); check(err);

      err = c.Insert( &Route{
	 byrouteid[rtid][1], 
	 intid,
	 intnumber,
	 false,
      } ); check(err);

   }

   for _, rtid := range ascendrouteids[index2834:] {

      intid, err := strconv.Atoi(byrouteid[rtid][2]); check(err)
      intnumber, err := strconv.Atoi(byrouteid[rtid][0]); check(err);

      err = c.Insert( &Route{
	 byrouteid[rtid][1], 
	 intid,
	 intnumber,
	 true,
      } ); check(err);

   }
   /*
   // count of collection 
   count, err := c.Count(); check(err);
   fmt.Println("count: " , count);

   // retrieve the whole collection
   routes := []Route{};
   iter := c.Find(nil).Iter(); 
   err = iter.All(&routes); check(err);
   for _, route := range routes {
      fmt.Println(route.Number, route.Name, route.Id);
   }
   */


}
