---
title: api
subTitle: get arrival estimates
---

# Get Arrival Estimates
`https://fivepoints.herokuapp.com/v1/get-arrival-estimates`
This endpoint returns all events for a certain station and direction for a given period.
The caller specifies a `StartDate` and `EndDate`, which creates an `EVENT_TIME` period. Only events within this bound will be returned.

## Request
```protobuf
message GetArrivalEstimatesRequest {
    google.protobuf.Timestamp StartDate = 1;
    google.protobuf.Timestamp EndDate = 2;
    string Station = 3;
    string Destination = 4;
    string LastEvaluatedKey = 5;
}
```
| Field            | Description                                              | Required |
|------------------|----------------------------------------------------------|----------|
| StartDate        | Must be `RFC3339`, and same date as end date             | `True`   |
| EndDate          | Must be `RFC3339`, and same date as start date           | `True`   |
| Station          | MARTA Station whose event times you are querying for     | `True`   |
| Destination      | The destination station of the line you are querying for | `True`   |
| LastEvaluatedKey | Provided in response, used to page through results       | `False`  |

### Example Request
```json
{
"StartDate": "2019-08-08T12:30:11+00:00",
"EndDate": "2019-08-08T12:33:11+00:00",
"Station": "AIRPORT STATION",
"Destination": "Airport"
}
```

## Response

```protobuf
message GetArrivalEstimatesResponse {
  repeated ArrivalEstimate ArrivalEstimates = 1;
  string LastEvaluatedKey = 2;
  int32 ResultLength = 3;
}

message ArrivalEstimate {
    string PrimaryKey = 1;
    string SortKey = 2;
    string Destination = 3;
    string Direction = 4;
    string EventTime = 5;
    string Line = 6;
    string NextArrival = 7;
    string Station = 8;
    string TrainID = 9;
    string WaitingSeconds = 10;
    string WaitingTime = 11;
    int64 TTL = 12;
}
```

### Expected Status Codes

## Authentication
This endpoint requires an `Authorization: Bearer jwt` header, with claim `fvp.Do` present.

If it does not exist, the caller will not be able to retrieve any information.
