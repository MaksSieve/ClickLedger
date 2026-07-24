package service

import (
	"clickledger/internal/model"
	"clickledger/internal/repository"
	"time"

	"gorm.io/gorm"
)

type DashboardService struct {
	analyticsRepo *repository.AnalyticsRepository
	linkRepo      *repository.LinkRepository
}

type UserDashboardData struct {
	Summary    Summary           `json:"summary"`
	TimeSeries []TimeSeriesEntry `json:"timeSeries"`
	TopSources []Source          `json:"topSources"`
	TopLinks   []LinkData        `json:"topLinks"`
}

type Summary struct {
	TotalClicks      ClickSummary `json:"totalClicks"`
	UniqueClicks     ClickSummary `json:"uniqueClicks"`
	UniqueClickRate  ClickSummary `json:"activeLinks"`
	ActiveLinks      ClickSummary `json:"uniqueClickRate"`
	HumanTrafficRate ClickSummary `json:"humanTrafficRate"`
}

type ClickSummary struct {
	Value         int `json:"value"`
	PreviousValue int `json:"previousValue"`
}

type TimeSeriesEntry struct {
	TimeStamp    string `json:"timeStamp"`
	Clicks       int    `json:"clicks"`
	UniqueClicks int    `json:"uniqueClicks"`
	HumanClicks  int    `json:"humanClicks"`
}

type Source struct {
	Key          string `json:"key"`
	Label        string `json:"label"`
	Clicks       int    `json:"clicks"`
	UniqueClicks int    `json:"uniqueClicks"`
	SharePercent int    `json:"sharepercent"`
}

type LinkData struct {
	Link         model.Link `json:"link"`
	Campaign     string     `json:"campaign"`
	Clicks       int        `json:"clicks"`
	UniqueClicks int        `json:"uniqueClicks"`
	SharePercent int        `json:"sharepercent"`
}

func CreateDashboardService(db *gorm.DB) *DashboardService {
	return &DashboardService{
		analyticsRepo: repository.CreateAnalyticsRepo(db),
		linkRepo:      repository.CreateLinkRepo(db),
	}
}

func (s *DashboardService) GetUserDashboard(from time.Time, to time.Time, granularity string) (*UserDashboardData, error) {

	fy, fm, fd := from.Date()

	var prevFrom, prevTo time.Time

	switch granularity {
	case "hour":
		{
			prevFrom = from.Add(-time.Hour)
			prevFrom = to.Add(-time.Hour)
		}
	case "day":
		{
			fy, fm, fd := from.Date()
			ty, tm, td := to.Date()

			prevFrom := time.Date()
		}
	}

	totalClicks, err := s.analyticsRepo.GetTotalClicks(from, to)
	prevTotalClicks, err := s.analyticsRepo.GetTotalClicks(from, to)

	if err != nil {
		return nil, err
	}

	dashboardData := &UserDashboardData{
		Summary: Summary{
			TotalClicks: totalClicks,
		},
	}

	return dashboardData, nil
}
