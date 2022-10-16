import  React,{useState, useEffect} from 'react';
import { styled } from '@mui/material/styles';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import Collapse from '@mui/material/Collapse';
import Avatar from '@mui/material/Avatar';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import { red } from '@mui/material/colors';
import FavoriteIcon from '@mui/icons-material/Favorite';
import ShareIcon from '@mui/icons-material/Share';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import MoreVertIcon from '@mui/icons-material/MoreVert';

import Grid2 from '@mui/material/Unstable_Grid2'
import KeyboardReview from '../components/ReviewCard'
import ReviewCard from '../components/ReviewCard'
import {useParams} from 'react-router-dom';

interface ExpandMoreProps extends IconButtonProps {
  expand: boolean;
}

interface Review {
  CreatedAt: string,
  UpdatedAt: string,
  DeletedAt: string,
  ID: number,
  keyboard_id: number,
  user_id: number,
  Stars: number,
  description: string,
  Likes: number
}

interface Keyboard {
  CreatedAt: string,
  DeletedAt: string, 
  ID: number,
  UpdatedAt: string,
  name: string,
  url_photo: string
}

export default function Reviews() {

  const [reviewsFetched, setReviewsFetched] = useState<Array<Review>> ([])
  const [keyBoardsFetched, setKeyoardsFetched] = useState<Array<Keyboard>>([])

  const {id} = useParams();

  useEffect(() => {
    fetch(process.env.REACT_APP_ENDPOINT+'getReviewsKeyboard/${id}')
      .then(res => res.json())
      .then( data => setReviewsFetched(data) )
    
    fetch(process.env.REACT_APP_ENDPOINT+'keyboards')
      .then(res => res.json())
      .then(dataKeyboards => {
        setKeyoardsFetched(dataKeyboards);
      })
  }, [])

  const handleUrl = () => {
    for (const keyboard of keyBoardsFetched){
      for(const review of reviewsFetched){
        if(review.ID == keyboard.ID){
          return keyboard.url_photo
        }
      }
    }

  }


  return (

    <Grid2 
        container 
        spacing={2} 
        paddingLeft={10} 
        paddingTop={5}>
        {reviewsFetched.map((item) => {

            return(
                <Grid2 xs={6} md={3}  >
                    <ReviewCard 
                       user_id = {item.user_id}
                       description = {item.description}
                       Stars ={item.Stars}
                       CreatedAt={item.CreatedAt}
                       url_photo = {handleUrl()}
                       userReview = {false}
                        />
                </Grid2>
            )
        })}         
    </Grid2>

  );
}
