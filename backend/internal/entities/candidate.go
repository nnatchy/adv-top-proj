package entities

type SearchCandidatesRequest struct {
	Position string `query:"position" validate:"required,oneof=year_one year_two year_three year_four representative"`
}

type SearchCandidatesResponse struct {
	StudentID          string `json:"student_id"`
	CandidateFirstName string `json:"candidate_first_name"`
	CandidateLastName  string `json:"candidate_last_name"`
	CandidatePosition  string `json:"candidate_position"`
	CandidateMajor     string `json:"candidate_major"`
	CandidateID        string `json:"id"`
}
