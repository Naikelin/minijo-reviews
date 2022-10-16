import * as React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Rating from '@mui/material/Rating';
import Grid2 from '@mui/material/Unstable_Grid2'
import { height } from '@mui/system';
import {Link} from 'react-router-dom';


interface Props {
  CreatedAt: string,
  DeletedAt: string, 
  ID: number,
  UpdatedAt: string,
  name: string,
  url_photo: string,
  stars: number,
}


function CardKeyboard(props: Props) {
  return (
    <Link to={`/reviews/${props.ID}`}>
      <Card 
        sx={{ maxWidth: 345, height: 245, display: 'flex', flexDirection: 'column', justifyContent: 'space-between', paddingTop: 1, cursor: 'pointer'}}>
        <CardMedia
          component="img" 
          height="140"
          image={ props.url_photo}
          alt="key1" />

        <CardContent sx={{height: 90, textOverflow: 'hidden'}}>
          <Typography gutterBottom variant="h5" component="div"> {props.name} </Typography>
          <Typography variant="body2" color="text.secondary"  > {''} </Typography>
        </CardContent>
        
        <CardContent>
          <Grid2 container display={'flex'} justifyContent={'space-between'}>
            <Rating name='read-only' value={props.stars} readOnly />
          </Grid2>
        </CardContent>
      </Card>
    </Link>
  );
}

export default CardKeyboard;