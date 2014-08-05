
package hsrgtfs


// download link address for gtfs zip file
var Zipwebaddress string = "http://googlehsrdocs.hamilton.ca/";

// local save path for downloaded gtfs zip file
var Zippath string = "~/Go/src/nad/hsrgtfs/zipfile/hsrgtfs.zip";


type SplitRoute struct { // custom data
   Outbound []int
   Inbound []int
}

type Route struct { // routes.txt 3 KB
   RouteLongname string // [0] route_long_name ex. York
   RouteId int // [5] route_id ex. 2903
   RouteShortname int // [8] route_short_name ex. 08

   DirectionId bool // custom data
                    // inferred by repeated routes
                    // 0  is outbound travel 
                    // 1 is return trip, inbound 
		    // might be E/W, N/S
		    // corresponds to custom Route.DirectionId?
}

type Trip struct { // trips.txt 1.1 MB
   RouteId int // [1] corresponds to Route.Id, route_id
   ShapeId int // [4] corresponds to shapes.shape_id
   TripId int // [6] corresponds to stop_times.trip_id
   DirectionId bool // [2] 
                    // 0 outbound travel 
                    // 1 return/inbound 
		    // corresponds to custom Route.DirectionId?
}

type Stoptime struct { // stop_times.txt 33 MB
   TripId int // [0] corresponds to tips
   ArrivalTime string // [1] arrival_time
   DepartureTime string // [2] departure_time
   StopId int // [3] stop_id ex. 2394
}

type Stop struct { // stops.txt 220KB
   StopId int // [0] corresponds to stop_times.stop_id
   StopeCode int // [1] ex. 2760
   StopName string // [2] ex BEACH
   StopLat float64 // [4] ex 54.36434
   StopLon float64 // [5] stop_lon
}

