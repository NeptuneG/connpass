package connpass

import "strconv"

type Query struct {
	EventID          []int    `json:"event_id"`
	KeywordAnd       []string `json:"keyword"`
	KeywordOr        []string `json:"keyword_or"`
	DateYearMonth    []int    `json:"ym"`
	DateYearMonthDay []int    `json:"ymd"`
	Nickname         []string `json:"nickname"`
	OwnerNickname    []string `json:"owner_nickname"`
	SeriesID         []int    `json:"series_id"`
	Start            int      `json:"start"`
	Order            int      `json:"order"`
	Count            int      `json:"count"`
	Format           string   `json:"format"`
}

func (q *Query) getQueryString() string {
	queryString := "?"

	if len(q.KeywordAnd) != 0 {
		for _, eventID := range q.EventID {
			queryString += "event_id=" + strconv.Itoa(eventID) + "&"
		}
	}

	if len(q.KeywordAnd) != 0 {
		for _, keywordAnd := range q.KeywordAnd {
			queryString += "keyword=" + keywordAnd + "&"
		}
	}

	if len(q.KeywordOr) != 0 {
		for _, keywordOr := range q.KeywordOr {
			queryString += "keyword_or=" + keywordOr + "&"
		}
	}

	if len(q.DateYearMonth) != 0 {
		for _, dateYearMonth := range q.DateYearMonth {
			queryString += "ym=" + strconv.Itoa(dateYearMonth) + "&"
		}
	}

	if len(q.DateYearMonthDay) != 0 {
		for _, dateYearMonthDay := range q.DateYearMonthDay {
			queryString += "ymd=" + strconv.Itoa(dateYearMonthDay) + "&"
		}
	}

	if len(q.Nickname) != 0 {
		for _, nickname := range q.Nickname {
			queryString += "nickname=" + nickname + "&"
		}
	}

	if len(q.OwnerNickname) != 0 {
		for _, ownerNickname := range q.OwnerNickname {
			queryString += "owner_nickname=" + ownerNickname + "&"
		}
	}

	if len(q.SeriesID) != 0 {
		for _, seriesID := range q.SeriesID {
			queryString += "series_id=" + strconv.Itoa(seriesID) + "&"
		}
	}

	if q.Start != 0 {
		queryString += "start=" + strconv.Itoa(q.Start) + "&"
	}

	if q.Order != 0 {
		queryString += "order=" + strconv.Itoa(q.Order) + "&"
	}

	if q.Count != 0 {
		queryString += "count=" + strconv.Itoa(q.Count) + "&"
	}

	if q.Format != "" && q.Format != "json" {
		queryString += "format=" + q.Format + "&"
	}

	return queryString[0 : len(queryString)-1]
}
