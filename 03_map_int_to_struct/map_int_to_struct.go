package main

import ("fmt"
		)

// Type of request
const (
		ISSUE_REQUEST = "ISSUE_REQUEST"
		BURN_REQUEST = "BURN_REQUEST"
)

// Reward points Request Status
const (
		ISSUED = "ISSUED"
		REQUESTED = "REQUESTED"
		PENDING_APPROVAL = "PENDING_APPROVAL"
		SETTLED_UP = "SETTLED_UP"
)

// Approval status from each entity
const (
		APPROVED_BY = "APPROVED_BY"
		APPROVED_DATE = "APPROVED_DATE"
		COMMENT = "COMMENT"
		APPROVEMENT_STATUS = "APPROVEMENT_STATUS"
		)
		
// Approval status
const (
		REJECTED = "REJECTED"
		ACCEPTED = "ACCEPTED"
		)

// Request
type Request struct{
	RequestID int  `json:"RequestID"`
	RequestType string `json:"RequestType"`
	RewardPoints uint `json:"RewardPoints"`
	RequestStatus string `json:"RequestStatus"`
	RequestedBy string `json:"RequestedBy"`
	RequestedDate string `json:"RequestedDate"`
	IssuedDate string `json:"IssuedDate"`
	identities []ApproveDetails `json:"identities"`
}

// Approval details
type ApproveDetails struct {
	ApprovedBy string `json:"ApprovedBy"`
	ApproveDate string `json:"ApproveDate"`
}

var requestID int = 0
var StateMap =  make(map[int]Request)


func main() {
	RequestRewardPoints(100)
}

func RequestRewardPoints(rewardPts uint) (int) {
	
	// Create new Request
	var RequestDetails Request

	RequestDetails.RequestID = requestID + 1
	RequestDetails.RequestType = ISSUE_REQUEST
	RequestDetails.RewardPoints = rewardPts
	RequestDetails.RequestStatus = REQUESTED
	RequestDetails.RequestedDate = "2018-09-20"
	RequestDetails.RequestedBy = "sunil"
	RequestDetails.identities = make([]ApproveDetails,0)

	// Call PutState
	StateMap[1] = RequestDetails
	approveRequest(1,"x","2018-08-20")
	approveRequest(1,"y","2018-08-20")
	approveRequest(1,"z","2018-08-20")
	
	return 0 // Raise an event
}


func approveRequest(requestID int, approvedBy string, approvedDate string) (int) {
	// GetState
	req := StateMap[requestID]
	req.identities = append(req.identities,ApproveDetails{approvedBy,approvedDate})
	StateMap[requestID] = req
	fmt.Println(req)
	return 0
}
