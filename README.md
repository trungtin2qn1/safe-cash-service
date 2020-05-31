# Safe cash service BE

## Overview:

For purpose of handling consumers info and videos

## Architect:

Architect of all over system:

![Architect](./doc/Architect.jpg)

## Database design:

![Database](./doc/Database.jpg)

## Local dep:

For porpose of education, you can check more detail in folder setup with docker and docker-compose for local dep

## Production env to use:

`url`: `http://35.198.235.51:6000/`

## APIs:

1. Register:

`url`: `/apis/register`

`method`: `POST`

`Header`: 

```json
{
    "Content-Type": "application/json"
}
```

`Request`:

```json
{
    "email": "trungtin2qn1@gmail.com",
	"password": "1234567",
	"store_name": "Tin Huynh 1"
}
```

`Response`:

```json
{
    "id": "1544916627202511873",
    "email": "trungtin2qn1@gmail.com",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJleHAiOjE1ODczOTI5NzAsImlzcyI6IlRpbiJ9.gK27j2jwyND7fi8uonD9u58X2jE9MiRAyz59FaQWoc1mcUS4kEniksKZxh4M3OfYnDinCBtjobfb0KTUEsj-7oQrixq2mf4p3hVTK8v1FCN50N6s0N2VDfgDFZ52rt5pnn6ugoBXjhUtdHUxVc8Scn7sLsSMmEyllYGw7Nog97eebBEx3lra5LlBQcGyv01nXi7pLvW4e4i4kq-0UKB9kVgfT4sbsFgiZiG7J5SLd11QX1T4PF0rpjidgy87Mm3R92aNfEZgu3eQcMW5yzvq3bHV6gEfxlGCY7zn8-3hESNxED4t6-Nsc61sDoEHQZhqkWCBE_KV5gmRvo1O_Ujz-siadxR30ToxJK7xl8MAhqjnpu_x9_bSrOAUEn3EX4PA7LYsWwYuA1ZVX-wejI4TyV7dJUz8-OBwyL7tLwQvU_IhuVGtrqU7cOacUOxSto4AD3R6RZJXmxjPWdWSpPQXChXUcb5jW-Dn2ZY9lYSn6f6OAV29btY8WKRTNi_wd9dBuvhzYSObYGTOzcqr2ExgXRO22S9_5mOqqc1i-tuhza0GOSI3_vViR8ZGWHmxNr-pJeVzXJ0PFKZ8IIKtQ7PVME1s_r1xDvUkKonhOqGPJalw5FxazmXHifQxrLb9LI1BkvZfl8fPwdPVD0Iazuk-6I1s6abHJANL126yTI91HaQ",
    "store_id": "1544916627202511873",
    "created_at": "2020-04-19T14:29:30.449653676Z",
    "updated_at": "2020-04-19T14:29:30.449653676Z"
}
```

2. Login: 

`url`: `/apis/login`

`method`: `POST`

`Header`:

```json
{
    "Content-Type": "application/json"
}
```

`Request`:

```json
{
	"email": "trungtin2qn1@gmail.com",
	"password": "1234567",
	"store_name": "Tin Huynh 1"
}
```

`Response`:

```json
{
    "id": "1544916627202511873",
    "email": "trungtin2qn1@gmail.com",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
    "store_id": "1544916627202511873",
    "created_at": "2020-04-19T14:29:30.449654Z",
    "updated_at": "2020-04-19T14:29:30.449654Z"
}
```

3. Upload media content:

`url`: `/apis/auth/media`

`method`: `POST`

`Header`: 

```json
{
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
    "Content-Type": "multipart/form-data; boundary=<calculated when request is sent>"
}
```

`Request`:

```json
{
    "file": binary file
}
```

`Response`:

```json
{
    "message": "Success",
}
```

4. Save notification token:

`url`: `/apis/auth/notification/token`

`Method`: `POST`

`Header`:

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
}
```

`Request`:

```json
{
	"token": "123"
}
```

`Response`:

```json
{
    "id": "1545233875968463873",
    "user_id": "1544916627202511873",
    "created_at": "2020-04-20T00:59:49.357387251Z",
    "updated_at": "2020-04-20T00:59:49.357387251Z"
}
```

5. Send notification:

`url`: `/apis/auth/notification`

`Method`: `POST`

`Header`: 

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
}
```

`Response`:

```json
{
    "message": "success"
}
```

6. Add user for a store:

`url`: `/apis/auth/register`

`Method`: `POST`

`Header`:

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
}
```

`Request`:

```json
{
    "email": string,
    "password": string
}
```

`Response`:

```json
{
    "id": "1544916627202511873",
    "email": "trungtin2qn1@gmail.com",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
    "store_id": "1544916627202511873",
    "created_at": "2020-04-19T14:29:30.449654Z",
    "updated_at": "2020-04-19T14:29:30.449654Z"
}
```