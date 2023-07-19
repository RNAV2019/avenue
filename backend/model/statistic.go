package model

func CreateStatistic(statistic Statistic) (*Statistic, error) {
	err := db.Create(&statistic).Error
	if err != nil {
		return nil, err
	}
	return &statistic, nil
}

func GetAggregateClicks(avenueID int) (int64, error) {
	var count int64
	err := db.Model(&Statistic{}).Where("avenue_id = ?", avenueID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetAllStatistics(avenueID int) ([]Statistic, error) {
	var statistics []Statistic
	err := db.Model(&Statistic{}).Where("avenue_id = ?", avenueID).Find(&statistics).Error
	if err != nil {
		return []Statistic{}, err
	}
	return statistics, nil
}
