/*
 * Copyright (c) 2023, Patrick Wilmes <patrick.wilmes@bit-lake.com>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */
package internal

const gitlabBaseUrl = "https://gitlab.com/api/v4/"

func UrlForMergeRequestPerProject(projectId string, accessToken string) string {
	return gitlabBaseUrl + "projects/" + projectId + "/merge_requests?state=opened&access_token=" + accessToken
}

func UrlForPostingANoteToAnMR(projectId string, mrId string, accessToken string) string {
	return gitlabBaseUrl + "projects/26177235/merge_requests/" + mrId + "/notes?access_token=" + accessToken
}

func UrlForDenyingAnMR(projectId string, mrId string, accessToken string) string {
	return gitlabBaseUrl + "projects/" + projectId + "/merge_requests/" + mrId + "?access_token=" + accessToken
}
