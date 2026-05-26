package jobqueue

const (
	// ResultStatusSuccess means that the job is successfully processed.
	ResultStatusSuccess = "success"

	// ResultStatusFailure means that the job is failed but it may be retried.
	ResultStatusFailure = "failure"

	// ResultStatusPermanentFailure means that the job is failed and
	// should never be retried.
	ResultStatusPermanentFailure = "permanent-failure"

	// ResultStatusInternalFailure means that the job is failed before
	// processing it in some internal reason.
	ResultStatusInternalFailure = "internal-failure"
)

// Result describes the result of a processed job.
type Result struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// IsSuccess returns if the job succeeded
func (rslt *Result) IsSuccess() bool { _ = "STUB: not implemented"; return false }

// IsFailure returns if the job is successfully processed or not.
func (rslt *Result) IsFailure() bool { _ = "STUB: not implemented"; return false }

// IsPermanentFailure returns if the job is permanently failed.
func (rslt *Result) IsPermanentFailure() bool { _ = "STUB: not implemented"; return false }

// IsFinished returns if the job can be retried or not.
func (rslt *Result) IsFinished() bool { _ = "STUB: not implemented"; return false }

// IsValid returns if the result status is valid or not.
func (rslt *Result) IsValid() bool { _ = "STUB: not implemented"; return false }
