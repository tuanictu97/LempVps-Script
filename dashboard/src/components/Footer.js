// Copyright 2023 tuanha1305
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

const Footer = () => {
    return (
      <div className='container'>
        <footer className='d-flex flex-wrap justify-content-between align-items-center py-3 my-4 border-top'>
          <div className='col-md-4 d-flex align-items-center'>
            <a
              href='/'
              className='mb-3 me-2 mb-md-0 text-muted text-decoration-none lh-1'
            >
              <svg className='bi' width={30} height={24}>
                <use xlinkHref='#bootstrap' />
              </svg>
            </a>
            <span className='text-muted'>© 2021 Company, Inc</span>
          </div>
          <ul className='nav col-md-4 justify-content-end list-unstyled d-flex'>
            <li className='ms-3'>
              <a className='text-muted' href='/'>
                <svg className='bi' width={24} height={24}>
                  <use xlinkHref='#twitter' />
                </svg>
              </a>
            </li>
            <li className='ms-3'>
              <a className='text-muted' href='/'>
                <svg className='bi' width={24} height={24}>
                  <use xlinkHref='#instagram' />
                </svg>
              </a>
            </li>
            <li className='ms-3'>
              <a className='text-muted' href='/'>
                <svg className='bi' width={24} height={24}>
                  <use xlinkHref='#facebook' />
                </svg>
              </a>
            </li>
          </ul>
        </footer>
      </div>
    )
  }
  
  export default Footer