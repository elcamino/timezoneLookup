package main
import (
	"time"
	"fmt"
	timezone "github.com/evanoberholster/timezoneLookup"
)

var tz timezone.TimezoneInterface

func main() {
	//tz = timezone.BoltdbStorage(timezone.WithSnappy, "timezone")
	tz = timezone.MemoryStorage(timezone.WithSnappy, "timezone")

	err := tz.LoadTimezones()
	if err != nil {
		fmt.Println(err)
	}

	querys := []timezone.Coord{
			{Y: 5.261417, X: -3.925778,}, // Abijan Airport
			{Y: -15.678889,X: 34.973889,}, // Blantyre Airport
			{Y: -12.65945, X: 18.25674,}, 
	    	{Y: 41.8976, X:-87.6205},
		    {Y: 47.6897, X: -122.4023},
		    {Y: 42.7235, X:-73.6931},
		    {Y: 42.5807, X:-83.0223},
		    {Y: 36.8381, X:-84.8500},
		    {Y: 40.1674, X:-85.3583},
		    {Y: 37.9643, X:-86.7453},
		    {Y: 38.6043, X:-90.2417},
		    {Y: 41.1591, X:-104.8261}, 
		    {Y: 35.1991, X:-111.6348}, 
		    {Y: 43.1432, X:-115.6750}, 
		    {Y: 47.5886, X:-122.3382}, 
		    {Y: 58.3168, X:-134.4397}, 
		    {Y: 21.4381, X:-158.0493}, 
		    {Y: 42.7000, X:-80.0000}, 
		    {Y: 51.0036, X:-114.0161}, 
		    {Y:-16.4965, X:-68.1702}, 
		    {Y:-31.9369, X:115.8453}, 
		    {Y: 42.0000, X:-87.5000}, 
	    	{Y: 41.8976, X:-87.6205},
		    {Y: 47.6897, X: -122.4023},
		    {Y: 42.7235, X:-73.6931},
		    {Y: 42.5807, X:-83.0223},
		    {Y: 36.8381, X:-84.8500},
		    {Y: 40.1674, X:-85.3583},
		    {Y: 37.9643, X:-86.7453},
		    {Y: 38.6043, X:-90.2417},
		    {Y: 41.1591, X:-104.8261}, 
		    {Y: 35.1991, X:-111.6348}, 
		    {Y: 43.1432, X:-115.6750}, 
		    {Y: 47.5886, X:-122.3382}, 
		    {Y: 58.3168, X:-134.4397}, 
		    {Y: 21.4381, X:-158.0493}, 
		    {Y: 42.7000, X:-80.0000}, 
		    {Y: 51.0036, X:-114.0161}, 
		    {Y:-16.4965, X:-68.1702}, 
		    {Y:-31.9369, X:115.8453}, 
		    {Y: 42.0000, X:-87.5000}, 
	    	{Y: 41.8976, X:-87.6205},
		    {Y: 47.6897, X: -122.4023},
		    {Y: 42.7235, X:-73.6931},
		    {Y: 42.5807, X:-83.0223},
		    {Y: 36.8381, X:-84.8500},
		    {Y: 40.1674, X:-85.3583},
		    {Y: 37.9643, X:-86.7453},
		    {Y: 38.6043, X:-90.2417},
		    {Y: 41.1591, X:-104.8261}, 
		    {Y: 35.1991, X:-111.6348}, 
		    {Y: 43.1432, X:-115.6750}, 
		    {Y: 47.5886, X:-122.3382}, 
		    {Y: 58.3168, X:-134.4397}, 
		    {Y: 21.4381, X:-158.0493}, 
		    {Y: 42.7000, X:-80.0000}, 
		    {Y: 51.0036, X:-114.0161}, 
		    {Y:-16.4965, X:-68.1702}, 
		    {Y:-31.9369, X:115.8453}, 
		    {Y: 42.0000, X:-87.5000}, 
		}
	
	for _, query := range querys {
		start := time.Now()
		res, err := tz.Query(query)
		if err != nil {
			fmt.Println(err)
		}
		elapsed := time.Since(start)
		fmt.Println("Query Result: ", res, " took: ", elapsed)
	}

	tz.Close()
}