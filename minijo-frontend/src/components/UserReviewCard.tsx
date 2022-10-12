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
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import FormControl from '@mui/material/FormControl';
import SaveIcon from '@mui/icons-material/Save';
import Stack from '@mui/material/Stack';
import Paper from '@mui/material/Paper';


interface ExpandMoreProps extends IconButtonProps {
  expand: boolean;
}

interface Props {
    id: number,
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


const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: 'center',
  color: theme.palette.text.secondary,
}));


export default function UserReviewCard(props: Props) {
  const [expanded, setExpanded] = React.useState(false);
  const [value, setValue] = React.useState<number | null>(props.Stars);
  const [newText, setText] = React.useState(props.description);

  const [editing, setEditing] = React.useState(false);
  

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  const handleDelete = () => {

    fetch(`http://localhost:5000/deleteReview`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        id: props.id
      })
    })
    .then(res => res.json())
    
  }

  const handleEditButton = () => {
    setEditing(true);
  }


  const handleEdit = () => {

    fetch(`http://localhost:5000/editReview`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        id: props.id,
        stars: 5,
        description: 'hola'
      })
    })
    .then(res => res.json())
    .then(data => console.log(data))
  }

  const handleFormSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    //get all the data from the form that was submitted
    const formData = new FormData(event.currentTarget);
    //create an object from the form data
    const data = Object.fromEntries(formData.entries());
    console.log(typeof data.stars);
    //send the data to the server
    fetch(`http://localhost:5000/editReview`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        id: props.id,
        stars: Number(data.stars),
        description: data.description
      })
    })
    .then(res => res.json())
    .then(data => console.log(data))
    //reset the form
    event.currentTarget.reset();
    //stop the page from reloading
    event.preventDefault();
    setEditing(false);

  }

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
        subheader={'Date: '+props.CreatedAt.substring(0,10)}>
      </CardHeader>
      
      <CardMedia
        component="img"
        height="194"
        image={props.url_photo}
        alt="Paella dish">
      </CardMedia>


      <CardContent>
        { 
        editing 
        ?
          <>
            <form onSubmit={handleFormSubmit}>
              <Stack spacing={2}>
                <TextField 
                  name="description" 
                  label="Description" 
                  variant="outlined"
                  value={newText}
                  defaultValue={newText} 
                  onChange={(newValue) => setText(newValue.target.value)} />
                <Rating
                  name="stars"
                  value={value}
                  onChange={(event, newValue) => {
                    setValue(newValue);
                  }}/>
                  <Button type="submit" variant="contained" startIcon={<SaveIcon />} >Save</Button>
              </Stack>
            </form>
          </>
        :
          <>
            <Typography variant="body2" color="text.secondary">
                {newText}
            </Typography>
            <Rating name="read-only" value={value} readOnly />
          </>
        }

      </CardContent>

      {
      editing
      ?
        <></>
      :
        <CardActions disableSpacing>
            <Button sx={{mr: 2 }} variant="contained" color="error" startIcon={<DeleteIcon />} onClick={handleDelete}>
              Delete
            </Button>
            <Button variant="contained"  startIcon={<EditIcon />} onClick={handleEditButton}>
              Edit
            </Button>
        </CardActions>
        }
     
    </Card>
  );
}
