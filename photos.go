
package main

type APIResponse struct { 
	Photos []struct { 
		ID int `json:"id"` 
		Sol int `json:"sol"` 
		Camera struct { 
			ID int `json:"id"` 
			Name string `json:"name"` 
			RoverID int `json:"rover_id"` 
			FullName string `json:"full_name"` 
		} `json:"camera"` 
		ImgSrc string `json:"img_src"` 
		EarthDate string `json:"earth_date"` 
		Rover struct { 
			ID int `json:"id"` 
			Name string `json:"name"` 
			LandingDate string `json:"landing_date"` 
			LaunchDate string `json:"launch_date"` 
			Status string `json:"status"` 
			MaxSol int `json:"max_sol"` 
			MaxDate string `json:"max_date"` 
			TotalPhotos int `json:"total_photos"` 
			Cameras []struct { 
				Name string `json:"name"` 
				FullName string `json:"full_name"` 
			} `json:"cameras"` 
		} `json:"rover"` 
	} `json:"photos"` 
}