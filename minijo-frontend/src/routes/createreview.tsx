import React, { useEffect, useState } from 'react';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import {useParams} from 'react-router-dom';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import { CardActionArea, Stack } from '@mui/material';
import TextField from '@mui/material/TextField';
import Rating from '@mui/material/Rating';
import Button from '@mui/material/Button';
import SaveIcon from '@mui/icons-material/Save';
import { Navigate } from "react-router-dom";

interface Keyboard {
    CreatedAt: string,
    DeletedAt: string, 
    ID: number,
    UpdatedAt: string,
    name: string,
    url_photo: string
}

function CreateReview() {

    const {iduser, idkeyboard} = useParams();
    const [keyboardData, setKeyboardData] = useState<Keyboard>({} as Keyboard);
    const [value, setValue] = React.useState<number | null>(0);
    const [newText, setText] = React.useState<string>('');
    const [finished, setFinished] = React.useState<boolean>(false);

    useEffect(() => {
        fetch(process.env.REACT_APP_ENDPOINT+`keyboards`)
        .then(res => res.json())
        .then((datos) => {
           datos.map((item: Keyboard) => {
                if(item.ID.toString() == idkeyboard){
                    setKeyboardData(item)
                }

           } )
        })
        
    }, [])


    const handleFormSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        //get all the data from the form that was submitted
        const formData = new FormData(event.currentTarget);
        //create an object from the form data
        const data = Object.fromEntries(formData.entries());
        console.log(typeof data.stars);
        //send the data to the server
        fetch(process.env.REACT_APP_ENDPOINT+`createReview`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            keyboard_id : idkeyboard,
            user_id: Number(iduser),
            description: data.description,
            stars: Number(data.stars)
          })
        })
        .then(res => res.json())
        .then(data => console.log(data))
        //reset the form
        event.currentTarget.reset();
        //stop the page from reloading
        event.preventDefault();
        setFinished(true);
      }

      if (finished) return <Navigate to='/' replace={true} />

    return (
        <>
        <Card sx={{ maxWidth: 345, m: 5}}>
        <CardActionArea>
          <CardMedia
            component="img"
            height="140"
            image={keyboardData.url_photo}
            alt="green iguana"
          />
          <CardContent>
            <Typography gutterBottom variant="h5" component="div">
              {keyboardData.name}
            </Typography>
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
          </CardContent>
        </CardActionArea>
      </Card>
      </>
        
    )
}

export default CreateReview;