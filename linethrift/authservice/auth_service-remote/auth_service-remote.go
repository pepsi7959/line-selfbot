// Code generated by Thrift Compiler (0.14.1). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/apache/thrift/lib/go/thrift"
	"talkservice"
	"authservice"
)

var _ = talkservice.GoUnusedProtection__
var _ = authservice.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  LoginResult loginZ(LoginRequest loginRequest)")
  fmt.Fprintln(os.Stderr, "  string confirmE2EELogin(string verifier, string deviceSecret)")
  fmt.Fprintln(os.Stderr, "  void respondE2EELoginRequest(string verifier, E2EEPublicKey publicKey, string encryptedKeyChain, string hashKeyChain, ErrorCode errorCode)")
  fmt.Fprintln(os.Stderr, "  string openAuthSession(AuthSessionRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse updatePassword(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  void logoutZ()")
  fmt.Fprintln(os.Stderr, "  string verifyQrcodeWithE2EE(string verifier, string pinCode, ErrorCode errorCode, E2EEPublicKey publicKey, string encryptedKeyChain, string hashKeyChain)")
  fmt.Fprintln(os.Stderr, "  RSAKey getAuthRSAKey(string authSessionId, IdentityProvider identityProvider)")
  fmt.Fprintln(os.Stderr, "  SecurityCenterResult issueTokenForAccountMigrationSettings(bool enforce)")
  fmt.Fprintln(os.Stderr, "  SetPasswordResponse setPassword(string authSessionId, EncryptedPassword encryptedPassword)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse confirmIdentifier(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse setIdentifier(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse setIdentifierAndPassword(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse updateIdentifier(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse resendIdentifierConfirmation(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse removeIdentifier(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := authservice.NewAuthServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "loginZ":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "LoginZ requires 1 args")
      flag.Usage()
    }
    arg62 := flag.Arg(1)
    mbTrans63 := thrift.NewTMemoryBufferLen(len(arg62))
    defer mbTrans63.Close()
    _, err64 := mbTrans63.WriteString(arg62)
    if err64 != nil {
      Usage()
      return
    }
    factory65 := thrift.NewTJSONProtocolFactory()
    jsProt66 := factory65.GetProtocol(mbTrans63)
    argvalue0 := authservice.NewLoginRequest()
    err67 := argvalue0.Read(context.Background(), jsProt66)
    if err67 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.LoginZ(context.Background(), value0))
    fmt.Print("\n")
    break
  case "confirmE2EELogin":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ConfirmE2EELogin requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.ConfirmE2EELogin(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "respondE2EELoginRequest":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "RespondE2EELoginRequest requires 5 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg71 := flag.Arg(2)
    mbTrans72 := thrift.NewTMemoryBufferLen(len(arg71))
    defer mbTrans72.Close()
    _, err73 := mbTrans72.WriteString(arg71)
    if err73 != nil {
      Usage()
      return
    }
    factory74 := thrift.NewTJSONProtocolFactory()
    jsProt75 := factory74.GetProtocol(mbTrans72)
    argvalue1 := talkservice.NewE2EEPublicKey()
    err76 := argvalue1.Read(context.Background(), jsProt75)
    if err76 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    tmp4, err := (strconv.Atoi(flag.Arg(5)))
    if err != nil {
      Usage()
     return
    }
    argvalue4 := authservice.ErrorCode(tmp4)
    value4 := argvalue4
    fmt.Print(client.RespondE2EELoginRequest(context.Background(), value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "openAuthSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "OpenAuthSession requires 1 args")
      flag.Usage()
    }
    arg79 := flag.Arg(1)
    mbTrans80 := thrift.NewTMemoryBufferLen(len(arg79))
    defer mbTrans80.Close()
    _, err81 := mbTrans80.WriteString(arg79)
    if err81 != nil {
      Usage()
      return
    }
    factory82 := thrift.NewTJSONProtocolFactory()
    jsProt83 := factory82.GetProtocol(mbTrans80)
    argvalue0 := authservice.NewAuthSessionRequest()
    err84 := argvalue0.Read(context.Background(), jsProt83)
    if err84 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.OpenAuthSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "updatePassword":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdatePassword requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg86 := flag.Arg(2)
    mbTrans87 := thrift.NewTMemoryBufferLen(len(arg86))
    defer mbTrans87.Close()
    _, err88 := mbTrans87.WriteString(arg86)
    if err88 != nil {
      Usage()
      return
    }
    factory89 := thrift.NewTJSONProtocolFactory()
    jsProt90 := factory89.GetProtocol(mbTrans87)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err91 := argvalue1.Read(context.Background(), jsProt90)
    if err91 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.UpdatePassword(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "logoutZ":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "LogoutZ requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.LogoutZ(context.Background()))
    fmt.Print("\n")
    break
  case "verifyQrcodeWithE2EE":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "VerifyQrcodeWithE2EE requires 6 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err := (strconv.Atoi(flag.Arg(3)))
    if err != nil {
      Usage()
     return
    }
    argvalue2 := authservice.ErrorCode(tmp2)
    value2 := argvalue2
    arg94 := flag.Arg(4)
    mbTrans95 := thrift.NewTMemoryBufferLen(len(arg94))
    defer mbTrans95.Close()
    _, err96 := mbTrans95.WriteString(arg94)
    if err96 != nil {
      Usage()
      return
    }
    factory97 := thrift.NewTJSONProtocolFactory()
    jsProt98 := factory97.GetProtocol(mbTrans95)
    argvalue3 := talkservice.NewE2EEPublicKey()
    err99 := argvalue3.Read(context.Background(), jsProt98)
    if err99 != nil {
      Usage()
      return
    }
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    fmt.Print(client.VerifyQrcodeWithE2EE(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "getAuthRSAKey":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetAuthRSAKey requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err := (strconv.Atoi(flag.Arg(2)))
    if err != nil {
      Usage()
     return
    }
    argvalue1 := authservice.IdentityProvider(tmp1)
    value1 := argvalue1
    fmt.Print(client.GetAuthRSAKey(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "issueTokenForAccountMigrationSettings":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "IssueTokenForAccountMigrationSettings requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1) == "true"
    value0 := argvalue0
    fmt.Print(client.IssueTokenForAccountMigrationSettings(context.Background(), value0))
    fmt.Print("\n")
    break
  case "setPassword":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SetPassword requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg105 := flag.Arg(2)
    mbTrans106 := thrift.NewTMemoryBufferLen(len(arg105))
    defer mbTrans106.Close()
    _, err107 := mbTrans106.WriteString(arg105)
    if err107 != nil {
      Usage()
      return
    }
    factory108 := thrift.NewTJSONProtocolFactory()
    jsProt109 := factory108.GetProtocol(mbTrans106)
    argvalue1 := authservice.NewEncryptedPassword()
    err110 := argvalue1.Read(context.Background(), jsProt109)
    if err110 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.SetPassword(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "confirmIdentifier":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ConfirmIdentifier requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg112 := flag.Arg(2)
    mbTrans113 := thrift.NewTMemoryBufferLen(len(arg112))
    defer mbTrans113.Close()
    _, err114 := mbTrans113.WriteString(arg112)
    if err114 != nil {
      Usage()
      return
    }
    factory115 := thrift.NewTJSONProtocolFactory()
    jsProt116 := factory115.GetProtocol(mbTrans113)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err117 := argvalue1.Read(context.Background(), jsProt116)
    if err117 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.ConfirmIdentifier(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "setIdentifier":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SetIdentifier requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg119 := flag.Arg(2)
    mbTrans120 := thrift.NewTMemoryBufferLen(len(arg119))
    defer mbTrans120.Close()
    _, err121 := mbTrans120.WriteString(arg119)
    if err121 != nil {
      Usage()
      return
    }
    factory122 := thrift.NewTJSONProtocolFactory()
    jsProt123 := factory122.GetProtocol(mbTrans120)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err124 := argvalue1.Read(context.Background(), jsProt123)
    if err124 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.SetIdentifier(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "setIdentifierAndPassword":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SetIdentifierAndPassword requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg126 := flag.Arg(2)
    mbTrans127 := thrift.NewTMemoryBufferLen(len(arg126))
    defer mbTrans127.Close()
    _, err128 := mbTrans127.WriteString(arg126)
    if err128 != nil {
      Usage()
      return
    }
    factory129 := thrift.NewTJSONProtocolFactory()
    jsProt130 := factory129.GetProtocol(mbTrans127)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err131 := argvalue1.Read(context.Background(), jsProt130)
    if err131 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.SetIdentifierAndPassword(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "updateIdentifier":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdateIdentifier requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg133 := flag.Arg(2)
    mbTrans134 := thrift.NewTMemoryBufferLen(len(arg133))
    defer mbTrans134.Close()
    _, err135 := mbTrans134.WriteString(arg133)
    if err135 != nil {
      Usage()
      return
    }
    factory136 := thrift.NewTJSONProtocolFactory()
    jsProt137 := factory136.GetProtocol(mbTrans134)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err138 := argvalue1.Read(context.Background(), jsProt137)
    if err138 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.UpdateIdentifier(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "resendIdentifierConfirmation":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ResendIdentifierConfirmation requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg140 := flag.Arg(2)
    mbTrans141 := thrift.NewTMemoryBufferLen(len(arg140))
    defer mbTrans141.Close()
    _, err142 := mbTrans141.WriteString(arg140)
    if err142 != nil {
      Usage()
      return
    }
    factory143 := thrift.NewTJSONProtocolFactory()
    jsProt144 := factory143.GetProtocol(mbTrans141)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err145 := argvalue1.Read(context.Background(), jsProt144)
    if err145 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.ResendIdentifierConfirmation(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "removeIdentifier":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "RemoveIdentifier requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg147 := flag.Arg(2)
    mbTrans148 := thrift.NewTMemoryBufferLen(len(arg147))
    defer mbTrans148.Close()
    _, err149 := mbTrans148.WriteString(arg147)
    if err149 != nil {
      Usage()
      return
    }
    factory150 := thrift.NewTJSONProtocolFactory()
    jsProt151 := factory150.GetProtocol(mbTrans148)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err152 := argvalue1.Read(context.Background(), jsProt151)
    if err152 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.RemoveIdentifier(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
