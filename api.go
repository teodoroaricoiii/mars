package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"os"
	"sync"
	"gopkg.in/urfave/cli.v1"
)

const (
    layoutISO = "2006-01-02"
	format_a = "01/02/06"
	format_b = "January 2, 2006"
	format_c = "Jan-02-06"
	format_d = "Jan-02-2006"
	URL = "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?api_key=DEMO_KEY&earth_date=" 
)

func convertDate(date string) (string, error) {

	t, err := time.Parse(format_a, date)
	if err == nil  {
                return t.Format(layoutISO), nil
	}
	
	t, err = time.Parse(format_b, date)
	if err == nil  {
                return t.Format(layoutISO), nil
        }

 
        t, err = time.Parse(format_c, date)
        if err == nil  {
                return t.Format(layoutISO), nil
        }
	
	t, err = time.Parse(format_d, date)
        if err == nil  {
                return t.Format(layoutISO), nil
        }     
 
        t, err = time.Parse(layoutISO, date)
        if err == nil  {
                return t.Format(layoutISO), nil
        }

        return "", errors.New("Invalid date format")	
}

func apiCall(date string)  {

	var response APIResponse

	// call the api
	cc := &http.Client{Timeout: time.Minute}
	res, err := cc.Get(URL + date)
	
	if err != nil  {
		fmt.Println(err)
		os.Exit(1)
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	//fmt.Printf("%s\n",b)
	//marshal and unmarshall the REST API result

	err = json.Unmarshal(b, &response)
	if err != nil  {
		fmt.Println(err)
		os.Exit(1)
	}


	// create a wait group
	var wg sync.WaitGroup


	//iterate response here
	for _, photo := range response.Photos  {
		wg.Add(1)
		go func(filename string)  {
			download(filename)
			wg.Done()
		}(photo.ImgSrc)
	}
	wg.Wait()
}

func main()  {
	//initialize command line
	app := cli.NewApp()
	app.Name = "mars-rover-api-client"
	app.Usage = "Download images using Mars Rover API"

	//create command line flags
	app.Flags = []cli.Flag{
		cli.StringFlag  {
			Name:	"date,d",
			Usage: "program --date [yyyy-mm-dd]",
		},
	}
	
	//use the command line
	app.Action = func(c * cli.Context) error  {
		date := c.GlobalString("date")
		t,err := convertDate(date)
		if err != nil  {
			fmt.Printf("Invalid date format")	
		}
		
		//do the REST API call
		apiCall(t)
		return nil;
	}

	//run the application
	app.Run(os.Args)
}
