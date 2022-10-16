import React from 'react';
import { Grid, Typography, Button, Box } from '@mui/material';
import myteam from '../images/myteam.jpg';

const Home = () => {
  

  return (
    <Box sx={{ 
      width: '100%',
      display: 'flex',
      minHeight: '600px',
      alignItems: 'center',
      justifyContent: 'center',}}>
      <Grid container spacing={6} sx={{display: 'flex',
      alignItems: 'center',
      maxWidth: '1300px',
      padding: '50px',}}>
        <Grid item xs={12} md={7}>
          <Typography variant="h3" fontWeight={700} sx={{ paddingBottom: '15px'}}>
            Minijo Reviews!
          </Typography>
          <Typography variant="h6" sx={{  opacity: '0.4',
      paddingBottom: '30px',}}>
            Evaluate your favorite keyboards and help others to choose the best one for them.
          </Typography>
          
        </Grid>
        <Grid item xs={12} md={5}>
          <img src="https://www.zdnet.com/a/img/resize/2dd017e31c3a58ee5125349380bc837ba6e2c8c1/2022/09/02/a92f3f18-2edc-4863-b2d5-9423c1a60644/keychron-k8-2.jpg?auto=webp&fit=crop&height=360&width=640" alt="My Team"/>
        </Grid>
      </Grid>
    </Box>
  );
};

export default Home;