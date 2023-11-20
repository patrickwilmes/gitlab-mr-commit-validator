/*
 * Copyright (c) 2023, Patrick Wilmes <patrick.wilmes@bit-lake.com>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */
package internal

import (
	"strconv"
)

type MR struct {
	ID      int  `json:"id"`
	IID     int  `json:"iid"`
	IsDraft bool `json:"draft"`
}

type mrDeny struct {
	Id    int    `json:"id"`
	IId   int    `json:"merge_request_iid"`
	State string `json:"state_event"`
}

func FetchMRs(token string) (mrs []MR, err error) {
	err = GetJson(UrlForMergeRequestPerProject(ProjectId, token), &mrs)
	return mrs, err
}

func (mr *MR) DenyMR(accessToken string) error {
	note := mrDeny{
		Id:    26177235,
		IId:   mr.IID,
		State: "close",
	}
	return ExecuteWriteRequest(PutRequest, UrlForDenyingAnMR(ProjectId, strconv.Itoa(mr.IID), accessToken), note)
}
