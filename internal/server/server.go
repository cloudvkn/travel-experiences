package server

import (
    "github.com/micro/go-micro/v3/api"
    "github.com/micro/go-micro/v3/micro"
    "github.com/cloudvkn/travel-experiences/booking"
    "github.com/cloudvkn/travel-experiences/itinerary"
    "github.com/cloudvkn/travel-experiences/recommendations"
)

func Run() {
    srv := micro.NewService(
        micro.Name("com.example.travelapp"),
        micro.Version("latest"),
    )

    srv.Init()

    bookingService := booking.NewBookingService()
    bookingHandler := booking.NewBookingHandler(bookingService)

    itineraryService := itinerary.NewItineraryService()
    itineraryHandler := itinerary.NewItineraryHandler(itineraryService)

    recommendationsService := recommendations.NewRecommendationsService()
    recommendationsHandler := recommendations.NewRecommendationsHandler(recommendationsService)

    apiHandler := api.NewHandler()
    apiHandler.HandleFunc("/booking", bookingHandler.Book)
    apiHandler.HandleFunc("/itinerary", itineraryHandler.PlanItinerary)
    apiHandler.HandleFunc("/recommendations", recommendationsHandler.GetRecommendations)

    srv.Handle(apiHandler)

    if err := srv.Run(); err != nil {
        panic(err)
    }
}
