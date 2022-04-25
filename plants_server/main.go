package main

import (
	h "GO_gRPC/handlers"
	models "GO_gRPC/models"
	plt "GO_gRPC/plants_service_grpc"
	"context"
	"encoding/json"
	"log"
	"net"

	"google.golang.org/grpc"
)

type PlantsServer struct {
	plt.UnimplementedPlantResolverServer
	dbHandler h.DBHandler
}

//Works
func (ps *PlantsServer) AddPlantProc(ctx context.Context, req *plt.AddPlantRequest) (resp *plt.AddPlantResponse, err error) {

	var data []byte
	var respErr []byte
	var plant = models.Plants{}

	json.Unmarshal(req.GetData(), &plant)

	if creationResult := ps.dbHandler.DB.Create(&plant); creationResult.Error != nil {
		data, _ = json.Marshal(map[string]any{"data": nil})
		respErr, _ = json.Marshal(map[string]any{"error": "Unanable to add plant"})

		err = creationResult.Error

		return &plt.AddPlantResponse{Data: data, Error: respErr}, err
	}

	data, _ = json.Marshal(map[string]any{"data": "Defined plant was correctly added"})
	respErr, _ = json.Marshal(map[string]any{"error": nil})

	return &plt.AddPlantResponse{Data: data, Error: respErr}, nil
}

//Works
func (ps *PlantsServer) GetPlantsProc(ctx context.Context, req *plt.GetPlantsRequest) (resp *plt.GetPlantsResponse, err error) {

	var data []byte
	var respErr []byte
	var plants = []models.Plants{}

	//ps.dbHandler.DB.Debug()

	if extractionResult := ps.dbHandler.DB.Find(&plants); extractionResult.Error != nil {
		data, _ = json.Marshal(map[string]any{"data": nil})
		respErr, _ = json.Marshal(map[string]any{"error": "Unanable to get plants"})

		err = extractionResult.Error

		return &plt.GetPlantsResponse{Data: data, Error: respErr}, err
	}

	data, _ = json.Marshal(map[string]any{"data": plants})
	respErr, _ = json.Marshal(map[string]any{"error": nil})

	return &plt.GetPlantsResponse{Data: data, Error: respErr}, nil
}

//Works
func (ps *PlantsServer) GetPlantByIdProc(ctx context.Context, req *plt.GetPlantByIdRequest) (resp *plt.GetPlantByIdResponse, err error) {

	var data []byte
	var respErr []byte
	var plant = models.Plants{}

	json.Unmarshal(req.GetData(), &plant)

	var id = plant.Id

	if extractionResult := ps.dbHandler.DB.First(&plant, id); extractionResult.Error != nil {
		data, _ = json.Marshal(map[string]any{"data": nil})
		respErr, _ = json.Marshal(map[string]any{"error": "Unanable to get defined plant"})

		err = extractionResult.Error

		return &plt.GetPlantByIdResponse{Data: data, Error: respErr}, err
	}

	data, _ = json.Marshal(map[string]any{"data": plant})
	respErr, _ = json.Marshal(map[string]any{"error": nil})

	return &plt.GetPlantByIdResponse{Data: data, Error: respErr}, nil
}

//Works
func (ps *PlantsServer) UpdatePlantByIdProc(ctx context.Context, req *plt.UpdatePlantByIdRequest) (resp *plt.UpdatePlantByIdResponse, err error) {

	var data []byte
	var respErr []byte
	var plant = models.Plants{}
	var updatedPlant = models.Plants{}

	json.Unmarshal(req.GetData(), &updatedPlant)

	var id = updatedPlant.Id

	if extractionResult := ps.dbHandler.DB.First(&plant, id); extractionResult.Error != nil {
		data, _ = json.Marshal(map[string]any{"data": nil})
		respErr, _ = json.Marshal(map[string]any{"error": "Unanable to update defined plant"})

		err = extractionResult.Error

		return &plt.UpdatePlantByIdResponse{Data: data, Error: respErr}, err
	}

	plant = updatedPlant

	ps.dbHandler.DB.Save(&plant)

	data, _ = json.Marshal(map[string]any{"data": "Defined plant was updated"})
	respErr, _ = json.Marshal(map[string]any{"error": nil})

	return &plt.UpdatePlantByIdResponse{Data: data, Error: respErr}, nil
}

//Works
func (ps *PlantsServer) DeletePlantByIdProc(ctx context.Context, req *plt.DeletePlantByIdRequest) (resp *plt.DeletePlantByIdResponse, err error) {

	var data []byte
	var respErr []byte
	var plant = models.Plants{}
	var deletedPlant = models.Plants{}

	json.Unmarshal(req.GetData(), &deletedPlant)

	var id = deletedPlant.Id

	if extractionResult := ps.dbHandler.DB.First(&plant, id); extractionResult.Error != nil {
		data, _ = json.Marshal(map[string]any{"data": nil})
		respErr, _ = json.Marshal(map[string]any{"error": "Unanable to delete defined plant"})

		err = extractionResult.Error

		return &plt.DeletePlantByIdResponse{Data: data, Error: respErr}, err
	}

	plant = deletedPlant

	ps.dbHandler.DB.Delete(&plant)

	data, _ = json.Marshal(map[string]any{"data": "Defined plant was deleted"})
	respErr, _ = json.Marshal(map[string]any{"error": nil})

	return &plt.DeletePlantByIdResponse{Data: data, Error: respErr}, nil
}

func main() {

	listener, err := net.Listen("tcp4", "localhost:8080")

	if err != nil {
		log.Fatalf("Unanable to listen port 8080: %v", err)
	}

	plantGRPCServer := grpc.NewServer()

	dbh := h.NewDBHandler()

	ps := PlantsServer{dbHandler: dbh}

	plt.RegisterPlantResolverServer(plantGRPCServer, &ps)

	log.Printf("Server started %v", listener.Addr())

	err = plantGRPCServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
