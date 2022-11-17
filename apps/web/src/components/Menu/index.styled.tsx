import styled from 'styled-components'

export const Nav = styled.nav`
  padding: 15px;
  border-bottom: 1px solid rgba(255, 62, 0, 0.1);
  font-weight: 300;

  .lang {
    float: right;
    svg {
      width: 32px;
      height: auto;
      cursor: pointer;
    }
  }

  ul {
    display: inline-block;
    padding: 0 !important;
    margin: 0 -32px 0 0 !important;
    list-style-type: none;

    li {
      display: inline-block;
      margin: auto 6px;
      * {
        vertical-align: middle;
      }
    }
  }

  a {
    text-decoration: none;
  }
`

export const Gitlab = styled.a`
  margin: 0 3px;

  svg {
    width: 32px !important;
    height: 32px;
  }
`
