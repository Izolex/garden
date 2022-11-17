import styled from 'styled-components'

export const Wrapper = styled.div`
  height: 100vh;
`

export const Main = styled.main`
  padding: 0 15px 40px 15px;
`

export const Container = styled.div`
  max-width: 960px;
  min-height: 100%;
  margin: 0 auto;
  text-align: center;
`

export const Footer = styled.footer`
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 40px;
  margin-top: -80px;
  text-align: center;

  svg {
    width: 50px;
    max-width: 20px;
    animation: pulse 1s ease infinite;
    fill: red;
  }

  @keyframes pulse {
    0% {
      transform: scale(1);
    }
    50% {
      transform: scale(1.3);
    }
    100% {
      transform: scale(1);
    }
  }
`
