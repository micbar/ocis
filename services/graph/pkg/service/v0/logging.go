package svc

import (
	"net/http"

	"github.com/owncloud/ocis/v2/ocis-pkg/log"
)

// NewLogging returns a service that logs messages.
func NewLogging(next Service, logger log.Logger) Service {
	return logging{
		next:   next,
		logger: logger,
	}
}

type logging struct {
	next   Service
	logger log.Logger
}

// ServeHTTP implements the Service interface.
func (l logging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.next.ServeHTTP(w, r)
}

// GetMe implements the Service interface.
func (l logging) GetMe(w http.ResponseWriter, r *http.Request) {
	l.next.GetMe(w, r)
}

// GetUsers implements the Service interface.
func (l logging) GetUsers(w http.ResponseWriter, r *http.Request) {
	l.next.GetUsers(w, r)
}

// GetUser implements the Service interface.
func (l logging) GetUser(w http.ResponseWriter, r *http.Request) {
	l.next.GetUser(w, r)
}

// PostUser implements the Service interface.
func (l logging) PostUser(w http.ResponseWriter, r *http.Request) {
	l.next.PostUser(w, r)
}

// DeleteUser implements the Service interface.
func (l logging) DeleteUser(w http.ResponseWriter, r *http.Request) {
	l.next.DeleteUser(w, r)
}

// PatchUser implements the Service interface.
func (l logging) PatchUser(w http.ResponseWriter, r *http.Request) {
	l.next.PatchUser(w, r)
}

// ChangeOwnPassword implements the Service interface.
func (l logging) ChangeOwnPassword(w http.ResponseWriter, r *http.Request) {
	l.next.ChangeOwnPassword(w, r)
}

// GetGroups implements the Service interface.
func (l logging) GetGroups(w http.ResponseWriter, r *http.Request) {
	l.next.GetGroups(w, r)
}

// GetGroup implements the Service interface.
func (l logging) GetGroup(w http.ResponseWriter, r *http.Request) {
	l.next.GetGroup(w, r)
}

// PostGroup implements the Service interface.
func (l logging) PostGroup(w http.ResponseWriter, r *http.Request) {
	l.next.PostGroup(w, r)
}

// PatchGroup implements the Service interface.
func (l logging) PatchGroup(w http.ResponseWriter, r *http.Request) {
	l.next.PatchGroup(w, r)
}

// DeleteGroup implements the Service interface.
func (l logging) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	l.next.DeleteGroup(w, r)
}

// GetGroupMembers implements the Service interface.
func (l logging) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	l.next.GetGroupMembers(w, r)
}

// PostGroupMember implements the Service interface.
func (l logging) PostGroupMember(w http.ResponseWriter, r *http.Request) {
	l.next.PostGroupMember(w, r)
}

// DeleteGroupMember implements the Service interface.
func (l logging) DeleteGroupMember(w http.ResponseWriter, r *http.Request) {
	l.next.DeleteGroupMember(w, r)
}

// GetSchools implements the Service interface.
func (l logging) GetSchools(w http.ResponseWriter, r *http.Request) {
	l.next.GetSchools(w, r)
}

// GetSchool implements the Service interface.
func (l logging) GetSchool(w http.ResponseWriter, r *http.Request) {
	l.next.GetSchool(w, r)
}

// PostSchool implements the Service interface.
func (l logging) PostSchool(w http.ResponseWriter, r *http.Request) {
	l.next.PostSchool(w, r)
}

// PatchSchool implements the Service interface.
func (l logging) PatchSchool(w http.ResponseWriter, r *http.Request) {
	l.next.PatchSchool(w, r)
}

// DeleteSchool implements the Service interface.
func (l logging) DeleteSchool(w http.ResponseWriter, r *http.Request) {
	l.next.DeleteSchool(w, r)
}

// GetSchoolMembers implements the Service interface.
func (l logging) GetSchoolMembers(w http.ResponseWriter, r *http.Request) {
	l.next.GetSchoolMembers(w, r)
}

// PostSchoolMember implements the Service interface.
func (l logging) PostSchoolMember(w http.ResponseWriter, r *http.Request) {
	l.next.PostSchoolMember(w, r)
}

// DeleteSchoolMember implements the Service interface.
func (l logging) DeleteSchoolMember(w http.ResponseWriter, r *http.Request) {
	l.next.DeleteSchoolMember(w, r)
}

// GetEducationUsers implements the Service interface.
func (l logging) GetEducationUsers(w http.ResponseWriter, r *http.Request) {
	l.next.GetEducationUsers(w, r)
}

// GetEducationUser implements the Service interface.
func (l logging) GetEducationUser(w http.ResponseWriter, r *http.Request) {
	l.next.GetEducationUser(w, r)
}

// PostEducationUser implements the Service interface.
func (l logging) PostEducationUser(w http.ResponseWriter, r *http.Request) {
	l.next.PostEducationUser(w, r)
}

// DeleteEducationUser implements the Service interface.
func (l logging) DeleteEducationUser(w http.ResponseWriter, r *http.Request) {
	l.next.DeleteEducationUser(w, r)
}

// PatchEducationUser implements the Service interface.
func (l logging) PatchEducationUser(w http.ResponseWriter, r *http.Request) {
	l.next.PatchEducationUser(w, r)
}

// GetDrives implements the Service interface.
func (l logging) GetDrives(w http.ResponseWriter, r *http.Request) {
	l.next.GetDrives(w, r)
}

// GetSingleDrive implements the Service interface.
func (l logging) GetSingleDrive(w http.ResponseWriter, r *http.Request) {
	l.next.GetDrives(w, r)
}

// UpdateDrive implements the Service interface.
func (l logging) UpdateDrive(w http.ResponseWriter, r *http.Request) {
	l.next.GetDrives(w, r)
}

// DeleteDrive implements the Service interface.
func (l logging) DeleteDrive(w http.ResponseWriter, r *http.Request) {
	l.next.GetDrives(w, r)
}

// GetAllDrives implements the Service interface.
func (l logging) GetAllDrives(w http.ResponseWriter, r *http.Request) {
	l.next.GetAllDrives(w, r)
}

// CreateDrive implements the Service interface.
func (l logging) CreateDrive(w http.ResponseWriter, r *http.Request) {
	l.next.CreateDrive(w, r)
}

// GetRootDriveChildren implements the Service interface.
func (l logging) GetRootDriveChildren(w http.ResponseWriter, r *http.Request) {
	l.next.GetRootDriveChildren(w, r)
}

// GetTags implements the Service interface.
func (l logging) GetTags(w http.ResponseWriter, r *http.Request) {
	l.next.GetTags(w, r)
}

// AssignTags implements the Service interface.
func (l logging) AssignTags(w http.ResponseWriter, r *http.Request) {
	l.next.AssignTags(w, r)
}

// UnassignTags implements the Service interface.
func (l logging) UnassignTags(w http.ResponseWriter, r *http.Request) {
	l.next.UnassignTags(w, r)
}
