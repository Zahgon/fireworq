//go:generate go-assets-builder -p mysql -o assets.go ../../data/jobqueue/mysql

package mysql

import (
	"regexp"
	"text/template"

	"github.com/fireworq/fireworq/model"
)

func newTableName(definition *model.Queue) *tableName { _ = "STUB: not implemented"; return nil }

type tableName struct {
	JobQueue string
	Payload  string
	Failure  string
}

func (tn *tableName) makeQueries() *sqls { _ = "STUB: not implemented"; return nil }

func (tn *tableName) makeQuery(tmpl *template.Template) string {
	_ = "STUB: not implemented"
	return ""
}

// ignore error

type sqls struct {
	createJobqueue     string
	createFailure      string
	grab               string
	grabbed            string
	launch             string
	insertJob          string
	insertFailedJob    string
	deleteFailedJob    string
	deleteJob          string
	updateJob          string
	orphan             string
	recover            string
	inspectJob         string
	inspectJobs        string
	inspectJobsAsc     string
	failedJob          string
	failedJobs         string
	recentlyFailedJobs string
}

var (
	invalidTablenameChars  *regexp.Regexp
	tmplCreateJobqueue     *template.Template
	tmplCreateFailure      *template.Template
	tmplGrabJobs           *template.Template
	tmplGrabbedJobs        *template.Template
	tmplLaunchJobs         *template.Template
	tmplInsertJob          *template.Template
	tmplInsertFailedJob    *template.Template
	tmplDeleteFailedJob    *template.Template
	tmplDeleteJob          *template.Template
	tmplUpdateJob          *template.Template
	tmplOrphanJobs         *template.Template
	tmplRecoverJobs        *template.Template
	tmplInspectJob         *template.Template
	tmplInspectJobs        *template.Template
	tmplInspectJobsAsc     *template.Template
	tmplFailedJob          *template.Template
	tmplFailedJobs         *template.Template
	tmplRecentlyFailedJobs *template.Template
)

func mustLoadTemplate(name string) *template.Template { _ = "STUB: not implemented"; return nil }

func init() {
	invalidTablenameChars = regexp.MustCompile("[^0-9a-z_]")
	tmplCreateJobqueue = mustLoadTemplate("schema/job_queue")
	tmplCreateFailure = mustLoadTemplate("schema/job_failure")
	tmplGrabJobs = mustLoadTemplate("query/grab_jobs")
	tmplGrabbedJobs = mustLoadTemplate("query/grabbed_jobs")
	tmplLaunchJobs = mustLoadTemplate("query/launch_jobs")
	tmplInsertJob = mustLoadTemplate("query/insert_job")
	tmplInsertFailedJob = mustLoadTemplate("query/insert_failed_job")
	tmplDeleteFailedJob = mustLoadTemplate("query/delete_failed_job")
	tmplDeleteJob = mustLoadTemplate("query/delete_job")
	tmplUpdateJob = mustLoadTemplate("query/update_job")
	tmplOrphanJobs = mustLoadTemplate("query/orphan_jobs")
	tmplRecoverJobs = mustLoadTemplate("query/recover_jobs")
	tmplInspectJob = mustLoadTemplate("query/inspect_job")
	tmplInspectJobs = mustLoadTemplate("query/inspect_jobs")
	tmplInspectJobsAsc = mustLoadTemplate("query/inspect_jobs_asc")
	tmplFailedJob = mustLoadTemplate("query/failed_job")
	tmplFailedJobs = mustLoadTemplate("query/failed_jobs")
	tmplRecentlyFailedJobs = mustLoadTemplate("query/recently_failed_jobs")
}
