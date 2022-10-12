import * as React from 'react';
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
import ThumbUpIcon from '@mui/icons-material/ThumbUp';
import ThumbDown from '@mui/icons-material/ThumbDown';
import Rating from '@mui/material/Rating';
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import Button from '@mui/material/Button';

interface ExpandMoreProps extends IconButtonProps {
  expand: boolean;
}

interface Props {

    user_id: number,
    description: string,
    Stars: number,
    CreatedAt: string,
    url_photo: string | undefined,
    userReview: boolean
}

const ExpandMore = styled((props: ExpandMoreProps) => {
  const { expand, ...other } = props;
  return <IconButton {...other} />;
})(({ theme, expand }) => ({
  transform: !expand ? 'rotate(0deg)' : 'rotate(180deg)',
  marginLeft: 'auto',
  transition: theme.transitions.create('transform', {
    duration: theme.transitions.duration.shortest,
  }),
}));

export default function ReviewCard(props: Props) {
  const [expanded, setExpanded] = React.useState(false);

  

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  return (
    <Card sx={{ maxWidth: 345 }}>
      <CardHeader
        avatar={
          <Avatar src={"/broken-image.jpg"} />
        }
        action={
          <IconButton aria-label="settings">
            <MoreVertIcon />
          </IconButton>
        }
        title={'User: '+props.user_id}
        subheader={'Date: '+props.CreatedAt.substring(0,10)}
      />
      <CardMedia
        component="img"
        height="194"
        image={props.url_photo}
        alt="Paella dish"
      />
      <CardContent>
        <Typography variant="body2" color="text.secondary">
            {props.description}
        </Typography>
        <Rating name="read-only" value={props.Stars} readOnly />

      </CardContent>

      <CardActions disableSpacing>
        {!props.userReview
        ?
        <>
          <IconButton aria-label="add to favorites">
            <ThumbUpIcon />
          </IconButton>
          <IconButton aria-label="share">
            <ThumbDown />
          </IconButton>
        </>
        :
        < >
          <Button sx={{mr: 2 }} variant="contained" color="error" startIcon={<DeleteIcon />}>
            Delete
          </Button>
          <Button variant="contained"  startIcon={<EditIcon />}>
            Edit
          </Button>
        </>
      }
    
      </CardActions>
     
    </Card>
  );
}
