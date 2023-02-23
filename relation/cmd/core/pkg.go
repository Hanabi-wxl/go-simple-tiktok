package core

import (
	"relation/cmd/dal/db"
	"relation/cmd/dal/redis"
	"relation/cmd/model"
	"relation/cmd/pack"
	"relation/cmd/service"
	"relation/pkg/consts"
	"strconv"
	"sync"
)

type RelationService struct {
}

func getUserCountInfo(userId int64) (totalFav, workCount, starCount int64) {
	suid := strconv.Itoa(int(userId))
	var userWorkIds []int64

	// 判断works中是否存在checkUserId(userId)
	if exist := redis.CheckUserIdExistInWorks(suid); exist {
		userWorkIds = redis.GetVideoIdsInWorks(suid)
	} else {
		// 不存在则添加key
		redis.CreateUserIdInWorks(suid)
		redis.AddExpireInWorks(suid)
		videos := db.GetVideosByUserId(userId)
		for i := 0; i < len(videos); i++ {
			if ok := redis.AddVideoIdInWorks(suid, videos[i].VideoId); !ok {
				redis.DeleteUserIdInWorks(suid)
			} else {
				userWorkIds = append(userWorkIds, videos[i].VideoId)
			}
		}
	}
	workCount = int64(len(userWorkIds))

	// 判断stars中是否存在videoId
	for i := 0; i < len(userWorkIds); i++ {
		svid := strconv.Itoa(int(userWorkIds[i]))
		if existInStars := redis.CheckVideoIdExistInStars(svid); existInStars {
			// 获取点赞人数
			stars := redis.GetUserIdsInStars(svid)
			totalFav += int64(len(stars))
		} else {
			redis.CreateVideoIdInStars(svid)
			redis.AddExpireInStars(svid)
			// 获取点赞人数
			favorites := db.GetStarUserById(userWorkIds[i])
			for _, user := range favorites {
				if save := redis.AddUserIdInStars(svid, user.UserId); !save {
					// 出现异常删除Key
					redis.DeleteVideoIdInStars(svid)
				}
			}
			totalFav += int64(len(favorites))
		}
	}

	// 判断star中是否存在userId
	if existInStar := redis.CheckUserIdExistInStar(suid); existInStar {
		starVideos := redis.GetVideoIdsInStar(suid)
		starCount = int64(len(starVideos))
	} else {
		// 不存在Key则新增Key
		redis.CreateUserIdInStar(suid)
		// 添加过期时间
		redis.AddExpireInStar(suid)
		// 获取用户所有点赞信息 保存至redis
		favoriteList := db.GetFavoriteListByUserId(userId)
		for _, likeVideoId := range favoriteList {
			if save := redis.AddVideoIdInStar(suid, likeVideoId.VideoId); !save {
				// 如果redis保存出现问题则删掉该Key
				redis.DeleteUserIdInStar(suid)
			}
		}
		starCount = int64(len(favoriteList))
	}
	return totalFav, workCount, starCount
}

func addToFollowList(index int, fid, uid int64, userList *[]*service.User, wg *sync.WaitGroup) {
	defer wg.Done()
	infoById := db.GetUserInfoById(fid)
	followInfo := db.GetUserFollowInfo(fid, uid)
	totalFav, workCount, starCount := getUserCountInfo(fid)
	(*userList)[index] = pack.BuildAuthor(infoById, followInfo, fid, totalFav, workCount, starCount)
}
func addToFollowerList(index int, uid int64, follower *model.Follow, authorInfos *[]*service.User, wg *sync.WaitGroup) {
	defer wg.Done()
	followerId := follower.FollowerId
	userinfo := db.GetUserInfoById(followerId)
	followerInfo := db.GetUserFollowInfo(followerId, uid)
	totalFav, workCount, starCount := getUserCountInfo(followerId)
	(*authorInfos)[index] = pack.BuildAuthor(userinfo, followerInfo, followerId, totalFav, workCount, starCount)
}

func addToFriendList(index int, friend model.Follow, uid int64, friendInfos *[]*service.FriendUser, wg *sync.WaitGroup) {
	defer wg.Done()
	var friendInfo service.FriendUser
	followerId := friend.FollowerId
	userinfo := db.GetUserInfoById(followerId)
	followerInfo := db.GetUserFollowInfo(followerId, uid)
	message := db.GetLastMessage(followerId, uid)
	messageType := message.CheckMessageType(uid)
	totalFav, workCount, starCount := getUserCountInfo(followerId)
	friendInfo.Id = &userinfo.UserId
	friendInfo.Name = &userinfo.Name
	friendInfo.Message = &message.Content
	friendInfo.Avatar = &userinfo.Avatar
	friendInfo.MsgType = &messageType
	friendInfo.FollowCount = &followerInfo.FollowCount
	friendInfo.FollowerCount = &followerInfo.FollowerCount
	friendInfo.IsFollow = &followerInfo.IsFollow
	friendInfo.Signature = &userinfo.Signature
	friendInfo.FavoriteCount = &starCount
	friendInfo.WorkCount = &workCount
	friendInfo.TotalFavorited = &totalFav
	url := consts.BackgroundImgUrl
	friendInfo.BackgroundImage = &url
	(*friendInfos)[index] = &friendInfo
}
