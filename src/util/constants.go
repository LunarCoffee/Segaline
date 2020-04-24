package util

const ResponseWriterBufferSize = 4_096
const MaxResponseSizeBeforeEncoding = 32_768

const HighestSupportedHttpVersion = HttpVersion11

const ContentLengthHeader = "Content-Length"
const ContentTypeHeader = "Content-Type"
const TransferEncodingHeader = "Transfer-Encoding"
