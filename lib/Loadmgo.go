
package hsrgtfs 

import (
   "archive/zip"
)

func Loadmgo(){

   // Open a zip archive for reading.
   r, err := zip.OpenReader(Zippath); check(err); defer r.Close();

   // Iterate files in the archive,
   for _, f := range r.File {

      switch f.Name {

	 case "routes.txt" : loadmgoroutes(f);
	 case "trips.txt" : loadmgotrips(f);
	 case "shapes.txt" : //todo
	 case "stop_times.txt" : loadmgostoptimes(f);
	 case "stops.txt" : loadmgostops(f);

	 

	 

      }
   }





}
