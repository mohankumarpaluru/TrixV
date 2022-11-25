// common variables
let configMaxUploadSize = window['MAX_UPLOAD_SIZE'];
let iBytesUploaded = 0
let iBytesTotal = 0
let iPreviousBytesLoaded = 0
let iMaxFilesize = configMaxUploadSize ? (+configMaxUploadSize) : (100 * 1024 * 1024); // 100MB
let timer = 0
let uploadInProgress = 'n/a'
let isProcessing = false
let file = null

/* CACHED ELEMENTS */

const uploadForm = document.getElementById('upload-form')
const videoInput = document.getElementById('video-input')
const targetLibraryPath = document.getElementById('target-library-path')
const videoTitle = document.getElementById('video-title')
const videoDescription = document.getElementById('video-description')
const uploadMessageLabel = document.getElementById('upload-message')
const uploadFileContainer = document.getElementById('upload-file')
const uploadFilenameLabel = document.getElementById('upload-filename')
const uploadButtonWrapper = document.getElementById('upload-button-wrapper')
const uploadButton = document.getElementById('upload-button')
const uploadProgressContainer = document.getElementById('upload-progress-container')
const uploadProgressBar = document.getElementById('upload-progress')
const uploadProgressLabel = document.getElementById('upload-progress-label')
const uploadStopped = document.getElementById('upload-stopped')
const uploadStarted = document.getElementById('upload-started')

/* HELPERS */

const setProgress = (_progress) => {
    uploadProgressContainer.style.display = _progress > 0 ? 'flex' : 'none'
    uploadProgressBar.style.width = `${_progress}%`
    uploadProgressLabel.innerText =  _progress >= 15 ? `${_progress}%` : ''
}

const setMessage = (_message, isError) => {
    uploadMessageLabel.style.display = _message ? 'block' : 'none'
    uploadMessageLabel.innerHTML = _message
    if (isError) {
        uploadMessageLabel.classList.add('error')
    } else {
        uploadMessageLabel.classList.remove('error')
    }
}

const setUploadState = (_uploadInProgress) => {
    uploadInProgress = _uploadInProgress

    uploadStarted.style.display = _uploadInProgress ? 'inline-block' : 'none'
    uploadStopped.style.display = _uploadInProgress ? 'none' : 'block'

    if (_uploadInProgress) {
        uploadButton.classList.add('transparent')
        timer = setInterval(doInnerUpdates, 300)
    } else {
        uploadButton.classList.remove('transparent')
        clearInterval(timer)
    }
}

const secondsToTime = (secs) => {
    let hr = Math.floor(secs / 3600)
    let min = Math.floor((secs - (hr * 3600)) / 60)
    let sec = Math.floor(secs - (hr * 3600) - (min * 60))
    if (hr < 10) hr = `0${hr}`
    if (min < 10) min = `0${min}`
    if (sec < 10) sec = `0${sec}`
    if (hr) hr = '00'
    return `${hr}:${min}:${sec}`
}

const bytesToSize = (bytes) => {
    const sizes = ['Bytes', 'KB', 'MB']
    if (bytes == 0) return 'n/a'
    const i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)))
    return (bytes / Math.pow(1024, i)).toFixed(1) + ' ' + sizes[i]
}

const determineDragAndDropCapable = () => {
    const div = document.createElement('div')
    return (('draggable' in div)
        || ('ondragstart' in div && 'ondrop' in div))
        && 'FormData' in window
        && 'FileReader' in window
}

/* MAIN */

document.addEventListener('DOMContentLoaded', () => {
    ['drag', 'dragstart', 'dragend', 'dragover', 'dragenter', 'dragleave']
        .forEach((evt) => {
            uploadForm.addEventListener(evt, (e) => {
                e.preventDefault()
                e.stopPropagation()
            })
        })

    uploadForm.addEventListener('drop', (e) => {
        console.log(111)
        e.preventDefault()
        e.stopPropagation()

        const _file = e.dataTransfer.files[0]
        if (!_file) return false
        fileSelected(_file)

        return false
    })
}, false)

const labelClicked = (e) => {
    if (uploadInProgress === true || file != null) {
        e.preventDefault()
        return false
    }
}

const fileSelected = (_file) => {
    const fileObj = _file || videoInput.files[0]
    if (_file) videoInput.value = ''

    if (fileObj) file = fileObj
    if (!file) return

    if (file.size > iMaxFilesize) {
        setMessage('Your file is too big. It can\'t be accepted. Please select a smaller file.', true)
        return
    }

    setMessage('')
    setProgress(0)
    
    const filename = file.name.length <= 20 ? file.name : `${file.name.substring(0, 14)}...${file.name.substring(file.name.length - 3)}`
    uploadFilenameLabel.innerText = filename
    uploadFileContainer.style.display = 'flex'
    
    uploadButtonWrapper.style.display = 'block'
    setUploadState(false)
}

const removeFile = (e, keepMessage) => {
    if (e) e.preventDefault()
    if (uploadInProgress === true) return

    uploadFileContainer.style.display = 'none'
    uploadButtonWrapper.style.display = 'none'
    videoInput.value = ''
    file = null
    if (!keepMessage) setMessage('No file selected')

    return false
}

const startUploading = () => {
    if (uploadInProgress === true) return
    if (!file) return

    isProcessing = false
    iPreviousBytesLoaded = 0
    setMessage('')
    setProgress(0)
    setUploadState(true)

    const formData = new FormData()
    formData.append('video_file', file)
    formData.append('target_library_path', targetLibraryPath.value)
    formData.append('video_title', videoTitle.value)
    formData.append('video_description', videoDescription.value)
    const xhr = new XMLHttpRequest()

    xhr.upload.addEventListener('progress', uploadProgress, false)
    xhr.addEventListener('load', uploadFinish, false)
    xhr.addEventListener('error', uploadError, false)
    xhr.addEventListener('abort', uploadAbort, false)

    xhr.open('POST', '/upload')
    xhr.send(formData)

    timer = setInterval(doInnerUpdates, 300)
}

const doInnerUpdates = () => { // we will use this function to display upload speed
    if (isProcessing) {
        clearInterval(timer)
        return
    }

    let iDiff = iBytesUploaded - iPreviousBytesLoaded
    // if nothing new loaded - exit
    if (iDiff == 0)
        return
    iPreviousBytesLoaded = iBytesUploaded
    iDiff = iDiff * 2
    const iBytesRem = iBytesTotal - iPreviousBytesLoaded
    const secondsRemaining = iBytesRem / iDiff
    // update speed info
    let iSpeed = iDiff.toString() + 'B/s'
    if (iDiff > 1024 * 1024) {
        iSpeed = (Math.round(iDiff * 100/(1024*1024))/100).toString() + 'MB/s'
    } else if (iDiff > 1024) {
        iSpeed =  (Math.round(iDiff * 100/1024)/100).toString() + 'KB/s'
    }

    const speedMessage = `${iSpeed} | ${secondsToTime(secondsRemaining)}`
    setMessage(speedMessage)
}

function uploadProgress(e) { // upload process in progress
    if (e.lengthComputable) {
        iBytesUploaded = e.loaded
        iBytesTotal = e.total

        const iPercentComplete = Math.round(iBytesUploaded / iBytesTotal * 100)
        setProgress(iPercentComplete)
        if (iPercentComplete === 100) {
            isProcessing = true
            setMessage('Processing video... please wait')
        }
    } else {
        setMessage('Unable to compute progress.')
    }
}

const uploadFinish = (e) => { // upload successfully finished
    const message = e.target.responseText
    const isSuccess = e.target.status < 400

    setProgress(isSuccess ? 100 : 0)
    setMessage(message, !isSuccess)
    setUploadState(false)
    if (isSuccess) removeFile(null, true)
}

const uploadError = () => { // upload error
    setMessage('An error occurred while uploading the file.', true)
    setProgress(0)
    setUploadState(false)
}

const uploadAbort = () => { // upload abort
    setMessage('The upload has been canceled by the user or the browser dropped the connection.', true)
    setProgress(0)
    setUploadState(false)
}
