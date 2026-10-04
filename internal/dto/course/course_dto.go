package course

type (
	//submit course from teacher
    SubmitCourseRequest struct {
      Title string `form:"title" validate:"required,max=200"`
    }
    
	//publish from admin
	PublishCourseRequest struct {
       SubmissionID int    `form:"submission_id" validate:"required"`
       YoutubeURL   string `form:"youtube_url" validate:"required,url"`
       YoutubeID    string `form:"youtube_id"`
    }

	
	//update
	UpdateCourseRequest struct {
		Title          string `form:"title" validate:"required"`
		YouTubeID     int `form:"file_id" validate:"required"`
	}
	
	//delete
	DeleteCourseResponse struct {
		Message string
	}
	//show all

	ShowCoursesResponse struct {
	    ID        int
	    Title     string
	    UserName string
        UserID  int
        YoutubeID  string
        YoutubeUrl string
    }  
)
