package api

import (
	"encoding/json"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/model/cms"
	"strconv"
	"time"
)

func (s *Service) GetSubCateByPid(lang string, pId int) (rows []*cms.NewsCategory, err error) {
	cacheKey := "GetSubCateByPid" + lang + strconv.Itoa(pId)
	if data := s.rep.BookCache().Get(cacheKey); data.Err() == nil {
		logger.Logger.Debug("load data from cache")
		err = json.Unmarshal([]byte(data.Val()), &rows)
		if err != nil {
			logger.Logger.Error(err.Error())
			return
		}
		return
	}

	var values interface{}
	values, err, _ = singleRequest.Do(cacheKey, func() (ret interface{}, err error) {
		rows, err = s.rep.GetSubCateByPid(lang, pId)
		for _, row := range rows {
			news, err := s.rep.GetCateNewsByPid(row.Id)
			if err == nil && len(news) > 0 {
				row.News = news
			}
		}
		logger.Logger.Debug("load data from db")
		return rows, err
	})
	rows = values.([]*cms.NewsCategory)
	var jsonData []byte
	jsonData, err = json.Marshal(rows)
	if err != nil {
		logger.Logger.Error(err.Error())
	}
	s.rep.BookCache().Set(cacheKey, string(jsonData), time.Duration(60)*time.Second)
	return
}

func (s *Service) GetCateNews(cateId, page int) (rows []cms.News, count int, err error) {
	cacheKey := "GetCateNews" + strconv.Itoa(cateId)
	if data := s.rep.BookCache().Get(cacheKey); data.Err() == nil {
		countC := s.rep.BookCache().Get(cacheKey + "count").Val()
		count, _ = strconv.Atoi(countC)
		logger.Logger.Debug("load data from cache")
		err = json.Unmarshal([]byte(data.Val()), &rows)
		if err != nil {
			logger.Logger.Error(err.Error())
			return
		}
		return
	}

	var values interface{}
	values, err, _ = singleRequest.Do(cacheKey, func() (ret interface{}, err error) {
		rows, count, err = s.rep.GetCateNews(cateId, page)
		data := map[string]interface{}{
			"count": count,
			"rows":  rows,
		}
		logger.Logger.Debug("load data from db")
		return data, err
	})
	data := values.(map[string]interface{})
	count = data["count"].(int)
	rows = data["rows"].([]cms.News)
	var jsonData []byte
	jsonData, err = json.Marshal(rows)
	if err != nil {
		logger.Logger.Error(err.Error())
	}
	s.rep.BookCache().Set(cacheKey, string(jsonData), time.Duration(60)*time.Second)
	s.rep.BookCache().Set(cacheKey+"count", count, time.Duration(60)*time.Second)
	return
}

func (s *Service) GetNewsDetail(id int) (row *cms.News, err error) {
	cacheKey := "GetNewsDetail" + strconv.Itoa(id)
	if data := s.rep.BookCache().Get(cacheKey); data.Err() == nil {
		logger.Logger.Debug("load data from cache")
		err = json.Unmarshal([]byte(data.Val()), &row)
		if err != nil {
			logger.Logger.Error(err.Error())
			return
		}
		return
	}

	var values interface{}
	values, err, _ = singleRequest.Do(cacheKey, func() (ret interface{}, err error) {
		row, err = s.rep.GetNewsDetail(id)
		logger.Logger.Debug("load data from db")
		return row, err
	})
	row = values.(*cms.News)
	var jsonData []byte
	jsonData, err = json.Marshal(row)
	if err != nil {
		logger.Logger.Error(err.Error())
	}
	s.rep.BookCache().Set(cacheKey, string(jsonData), time.Duration(60)*time.Second)
	return
}

func (s *Service) GetAboutUsNews(lang string, cId int) (rows []cms.News, err error) {
	cacheKey := "GetAboutUsNews" + lang + strconv.Itoa(cId)
	if data := s.rep.BookCache().Get(cacheKey); data.Err() == nil {
		logger.Logger.Debug("load data from cache")
		err = json.Unmarshal([]byte(data.Val()), &rows)
		if err != nil {
			logger.Logger.Error(err.Error())
			return
		}
		return
	}

	var values interface{}
	values, err, _ = singleRequest.Do(cacheKey, func() (ret interface{}, err error) {
		rows, err = s.rep.GetAboutUsNews(lang, cId)
		logger.Logger.Debug("load data from db")
		return rows, err
	})
	rows = values.([]cms.News)
	var jsonData []byte
	jsonData, err = json.Marshal(rows)
	if err != nil {
		logger.Logger.Error(err.Error())
	}
	s.rep.BookCache().Set(cacheKey, string(jsonData), time.Duration(60)*time.Second)
	return
}

func (s *Service) GetAdListByPosition(pId int) (rows []cms.AdMaterials, err error) {
	cacheKey := "GetAdListByPosition" + strconv.Itoa(pId)
	if data := s.rep.BookCache().Get(cacheKey); data.Err() == nil {
		logger.Logger.Debug("load data from cache")
		err = json.Unmarshal([]byte(data.Val()), &rows)
		if err != nil {
			logger.Logger.Error(err.Error())
			return
		}
		return
	}

	var values interface{}
	values, err, _ = singleRequest.Do(cacheKey, func() (ret interface{}, err error) {
		rows, err = s.rep.GetAdListByPosition(pId)
		logger.Logger.Debug("load data from db")
		return rows, err
	})
	rows = values.([]cms.AdMaterials)
	var jsonData []byte
	jsonData, err = json.Marshal(rows)
	if err != nil {
		logger.Logger.Error(err.Error())
	}
	s.rep.BookCache().Set(cacheKey, string(jsonData), time.Duration(60)*time.Second)
	return
}
