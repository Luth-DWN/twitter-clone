import { Tweet } from '../../tweets/types'
import { ProfileStatus, ProfileDetails } from '../types'

export interface State {
  status?: ProfileStatus
  profileDetails: ProfileDetails
  profileTweets: Tweet[]
}

export const state: State = {
  profileDetails: {
    id: 0,
    name: '',
    handle: '',
    bio: '',
    location: '',
    website: '',
    birthDate: '',
    followersCount: 0,
    followingsCount: 0,
    joinedAt: '',
  },
  profileTweets: [],
}
