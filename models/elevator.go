package models

import (
  "encoding/json"
  "fmt"
  "github.com/tunamsyar/go-elevator/utilities"
  "log"
  "os"
)

type Elevator struct {
  Queue []int `json:"queue"`
  LowZoneQueue []int `json:"low_zone_queue"`
  HighZoneQueue []int `json:"high_zone_queue"`
}

func NewElevator() *Elevator {
  return &Elevator{
    Queue: make([]int, 0),
    LowZoneQueue: make([]int, 0),
    HighZoneQueue: make([]int, 0),
  }
}

func (e *Elevator) AddFloor(floor int) {
  currentElevator := Elevator{}

  data, err := os.ReadFile("elevator_state.json")

  if err != nil {
    log.Printf("Error reading file: %v", err)
    return
  }

  err = json.Unmarshal(data, &currentElevator)

  if err != nil {
    log.Printf("Error unmarshal data: %v", err)
    return
  }

  currentElevator.Queue = utilities.SortAndRemoveDuplicates(append(currentElevator.Queue, floor))
  
  low, high := utilities.SplitArrayToLowHighZone(currentElevator.Queue)
  
  currentElevator.LowZoneQueue = low
  currentElevator.HighZoneQueue = high

  currentElevator.saveToFile()
}

func (e *Elevator) saveToFile() {
  fmt.Printf("ELEVATOR: %+v\n", e)
  data, err := json.MarshalIndent(e, "", "  ")

  fmt.Printf("DATA: %+v\n", string(data))

  if err != nil {
    log.Printf("Error marshaling elevator state: %v", err)
    return
  }
  err = os.WriteFile("elevator_state.json", data, 0644)

  if err != nil {
    log.Printf("Error writing elevator state to file: %v", err)
  }
}
