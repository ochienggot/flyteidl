// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/core/errors.proto

#ifndef PROTOBUF_flyteidl_2fcore_2ferrors_2eproto__INCLUDED
#define PROTOBUF_flyteidl_2fcore_2ferrors_2eproto__INCLUDED

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3005000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3005001 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)

namespace protobuf_flyteidl_2fcore_2ferrors_2eproto {
// Internal implementation detail -- do not use these members.
struct TableStruct {
  static const ::google::protobuf::internal::ParseTableField entries[];
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[];
  static const ::google::protobuf::internal::ParseTable schema[2];
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors();
void InitDefaultsContainerErrorImpl();
void InitDefaultsContainerError();
void InitDefaultsErrorDocumentImpl();
void InitDefaultsErrorDocument();
inline void InitDefaults() {
  InitDefaultsContainerError();
  InitDefaultsErrorDocument();
}
}  // namespace protobuf_flyteidl_2fcore_2ferrors_2eproto
namespace flyteidl {
namespace core {
class ContainerError;
class ContainerErrorDefaultTypeInternal;
extern ContainerErrorDefaultTypeInternal _ContainerError_default_instance_;
class ErrorDocument;
class ErrorDocumentDefaultTypeInternal;
extern ErrorDocumentDefaultTypeInternal _ErrorDocument_default_instance_;
}  // namespace core
}  // namespace flyteidl
namespace flyteidl {
namespace core {

enum ContainerError_Kind {
  ContainerError_Kind_NON_RECOVERABLE = 0,
  ContainerError_Kind_RECOVERABLE = 1,
  ContainerError_Kind_ContainerError_Kind_INT_MIN_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32min,
  ContainerError_Kind_ContainerError_Kind_INT_MAX_SENTINEL_DO_NOT_USE_ = ::google::protobuf::kint32max
};
bool ContainerError_Kind_IsValid(int value);
const ContainerError_Kind ContainerError_Kind_Kind_MIN = ContainerError_Kind_NON_RECOVERABLE;
const ContainerError_Kind ContainerError_Kind_Kind_MAX = ContainerError_Kind_RECOVERABLE;
const int ContainerError_Kind_Kind_ARRAYSIZE = ContainerError_Kind_Kind_MAX + 1;

const ::google::protobuf::EnumDescriptor* ContainerError_Kind_descriptor();
inline const ::std::string& ContainerError_Kind_Name(ContainerError_Kind value) {
  return ::google::protobuf::internal::NameOfEnum(
    ContainerError_Kind_descriptor(), value);
}
inline bool ContainerError_Kind_Parse(
    const ::std::string& name, ContainerError_Kind* value) {
  return ::google::protobuf::internal::ParseNamedEnum<ContainerError_Kind>(
    ContainerError_Kind_descriptor(), name, value);
}
// ===================================================================

class ContainerError : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.core.ContainerError) */ {
 public:
  ContainerError();
  virtual ~ContainerError();

  ContainerError(const ContainerError& from);

  inline ContainerError& operator=(const ContainerError& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  ContainerError(ContainerError&& from) noexcept
    : ContainerError() {
    *this = ::std::move(from);
  }

  inline ContainerError& operator=(ContainerError&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const ContainerError& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const ContainerError* internal_default_instance() {
    return reinterpret_cast<const ContainerError*>(
               &_ContainerError_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    0;

  void Swap(ContainerError* other);
  friend void swap(ContainerError& a, ContainerError& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline ContainerError* New() const PROTOBUF_FINAL { return New(NULL); }

  ContainerError* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  void CopyFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void MergeFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void CopyFrom(const ContainerError& from);
  void MergeFrom(const ContainerError& from);
  void Clear() PROTOBUF_FINAL;
  bool IsInitialized() const PROTOBUF_FINAL;

  size_t ByteSizeLong() const PROTOBUF_FINAL;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) PROTOBUF_FINAL;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const PROTOBUF_FINAL;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const PROTOBUF_FINAL;
  int GetCachedSize() const PROTOBUF_FINAL { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const PROTOBUF_FINAL;
  void InternalSwap(ContainerError* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const PROTOBUF_FINAL;

  // nested types ----------------------------------------------------

  typedef ContainerError_Kind Kind;
  static const Kind NON_RECOVERABLE =
    ContainerError_Kind_NON_RECOVERABLE;
  static const Kind RECOVERABLE =
    ContainerError_Kind_RECOVERABLE;
  static inline bool Kind_IsValid(int value) {
    return ContainerError_Kind_IsValid(value);
  }
  static const Kind Kind_MIN =
    ContainerError_Kind_Kind_MIN;
  static const Kind Kind_MAX =
    ContainerError_Kind_Kind_MAX;
  static const int Kind_ARRAYSIZE =
    ContainerError_Kind_Kind_ARRAYSIZE;
  static inline const ::google::protobuf::EnumDescriptor*
  Kind_descriptor() {
    return ContainerError_Kind_descriptor();
  }
  static inline const ::std::string& Kind_Name(Kind value) {
    return ContainerError_Kind_Name(value);
  }
  static inline bool Kind_Parse(const ::std::string& name,
      Kind* value) {
    return ContainerError_Kind_Parse(name, value);
  }

  // accessors -------------------------------------------------------

  // string code = 1;
  void clear_code();
  static const int kCodeFieldNumber = 1;
  const ::std::string& code() const;
  void set_code(const ::std::string& value);
  #if LANG_CXX11
  void set_code(::std::string&& value);
  #endif
  void set_code(const char* value);
  void set_code(const char* value, size_t size);
  ::std::string* mutable_code();
  ::std::string* release_code();
  void set_allocated_code(::std::string* code);

  // string message = 2;
  void clear_message();
  static const int kMessageFieldNumber = 2;
  const ::std::string& message() const;
  void set_message(const ::std::string& value);
  #if LANG_CXX11
  void set_message(::std::string&& value);
  #endif
  void set_message(const char* value);
  void set_message(const char* value, size_t size);
  ::std::string* mutable_message();
  ::std::string* release_message();
  void set_allocated_message(::std::string* message);

  // .flyteidl.core.ContainerError.Kind kind = 3;
  void clear_kind();
  static const int kKindFieldNumber = 3;
  ::flyteidl::core::ContainerError_Kind kind() const;
  void set_kind(::flyteidl::core::ContainerError_Kind value);

  // @@protoc_insertion_point(class_scope:flyteidl.core.ContainerError)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr code_;
  ::google::protobuf::internal::ArenaStringPtr message_;
  int kind_;
  mutable int _cached_size_;
  friend struct ::protobuf_flyteidl_2fcore_2ferrors_2eproto::TableStruct;
  friend void ::protobuf_flyteidl_2fcore_2ferrors_2eproto::InitDefaultsContainerErrorImpl();
};
// -------------------------------------------------------------------

class ErrorDocument : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.core.ErrorDocument) */ {
 public:
  ErrorDocument();
  virtual ~ErrorDocument();

  ErrorDocument(const ErrorDocument& from);

  inline ErrorDocument& operator=(const ErrorDocument& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  ErrorDocument(ErrorDocument&& from) noexcept
    : ErrorDocument() {
    *this = ::std::move(from);
  }

  inline ErrorDocument& operator=(ErrorDocument&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const ErrorDocument& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const ErrorDocument* internal_default_instance() {
    return reinterpret_cast<const ErrorDocument*>(
               &_ErrorDocument_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    1;

  void Swap(ErrorDocument* other);
  friend void swap(ErrorDocument& a, ErrorDocument& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline ErrorDocument* New() const PROTOBUF_FINAL { return New(NULL); }

  ErrorDocument* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  void CopyFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void MergeFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void CopyFrom(const ErrorDocument& from);
  void MergeFrom(const ErrorDocument& from);
  void Clear() PROTOBUF_FINAL;
  bool IsInitialized() const PROTOBUF_FINAL;

  size_t ByteSizeLong() const PROTOBUF_FINAL;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) PROTOBUF_FINAL;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const PROTOBUF_FINAL;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const PROTOBUF_FINAL;
  int GetCachedSize() const PROTOBUF_FINAL { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const PROTOBUF_FINAL;
  void InternalSwap(ErrorDocument* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const PROTOBUF_FINAL;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // .flyteidl.core.ContainerError error = 1;
  bool has_error() const;
  void clear_error();
  static const int kErrorFieldNumber = 1;
  const ::flyteidl::core::ContainerError& error() const;
  ::flyteidl::core::ContainerError* release_error();
  ::flyteidl::core::ContainerError* mutable_error();
  void set_allocated_error(::flyteidl::core::ContainerError* error);

  // @@protoc_insertion_point(class_scope:flyteidl.core.ErrorDocument)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::flyteidl::core::ContainerError* error_;
  mutable int _cached_size_;
  friend struct ::protobuf_flyteidl_2fcore_2ferrors_2eproto::TableStruct;
  friend void ::protobuf_flyteidl_2fcore_2ferrors_2eproto::InitDefaultsErrorDocumentImpl();
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// ContainerError

// string code = 1;
inline void ContainerError::clear_code() {
  code_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& ContainerError::code() const {
  // @@protoc_insertion_point(field_get:flyteidl.core.ContainerError.code)
  return code_.GetNoArena();
}
inline void ContainerError::set_code(const ::std::string& value) {
  
  code_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:flyteidl.core.ContainerError.code)
}
#if LANG_CXX11
inline void ContainerError::set_code(::std::string&& value) {
  
  code_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:flyteidl.core.ContainerError.code)
}
#endif
inline void ContainerError::set_code(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  code_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:flyteidl.core.ContainerError.code)
}
inline void ContainerError::set_code(const char* value, size_t size) {
  
  code_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:flyteidl.core.ContainerError.code)
}
inline ::std::string* ContainerError::mutable_code() {
  
  // @@protoc_insertion_point(field_mutable:flyteidl.core.ContainerError.code)
  return code_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* ContainerError::release_code() {
  // @@protoc_insertion_point(field_release:flyteidl.core.ContainerError.code)
  
  return code_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void ContainerError::set_allocated_code(::std::string* code) {
  if (code != NULL) {
    
  } else {
    
  }
  code_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), code);
  // @@protoc_insertion_point(field_set_allocated:flyteidl.core.ContainerError.code)
}

// string message = 2;
inline void ContainerError::clear_message() {
  message_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& ContainerError::message() const {
  // @@protoc_insertion_point(field_get:flyteidl.core.ContainerError.message)
  return message_.GetNoArena();
}
inline void ContainerError::set_message(const ::std::string& value) {
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:flyteidl.core.ContainerError.message)
}
#if LANG_CXX11
inline void ContainerError::set_message(::std::string&& value) {
  
  message_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:flyteidl.core.ContainerError.message)
}
#endif
inline void ContainerError::set_message(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:flyteidl.core.ContainerError.message)
}
inline void ContainerError::set_message(const char* value, size_t size) {
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:flyteidl.core.ContainerError.message)
}
inline ::std::string* ContainerError::mutable_message() {
  
  // @@protoc_insertion_point(field_mutable:flyteidl.core.ContainerError.message)
  return message_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* ContainerError::release_message() {
  // @@protoc_insertion_point(field_release:flyteidl.core.ContainerError.message)
  
  return message_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void ContainerError::set_allocated_message(::std::string* message) {
  if (message != NULL) {
    
  } else {
    
  }
  message_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), message);
  // @@protoc_insertion_point(field_set_allocated:flyteidl.core.ContainerError.message)
}

// .flyteidl.core.ContainerError.Kind kind = 3;
inline void ContainerError::clear_kind() {
  kind_ = 0;
}
inline ::flyteidl::core::ContainerError_Kind ContainerError::kind() const {
  // @@protoc_insertion_point(field_get:flyteidl.core.ContainerError.kind)
  return static_cast< ::flyteidl::core::ContainerError_Kind >(kind_);
}
inline void ContainerError::set_kind(::flyteidl::core::ContainerError_Kind value) {
  
  kind_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.core.ContainerError.kind)
}

// -------------------------------------------------------------------

// ErrorDocument

// .flyteidl.core.ContainerError error = 1;
inline bool ErrorDocument::has_error() const {
  return this != internal_default_instance() && error_ != NULL;
}
inline void ErrorDocument::clear_error() {
  if (GetArenaNoVirtual() == NULL && error_ != NULL) {
    delete error_;
  }
  error_ = NULL;
}
inline const ::flyteidl::core::ContainerError& ErrorDocument::error() const {
  const ::flyteidl::core::ContainerError* p = error_;
  // @@protoc_insertion_point(field_get:flyteidl.core.ErrorDocument.error)
  return p != NULL ? *p : *reinterpret_cast<const ::flyteidl::core::ContainerError*>(
      &::flyteidl::core::_ContainerError_default_instance_);
}
inline ::flyteidl::core::ContainerError* ErrorDocument::release_error() {
  // @@protoc_insertion_point(field_release:flyteidl.core.ErrorDocument.error)
  
  ::flyteidl::core::ContainerError* temp = error_;
  error_ = NULL;
  return temp;
}
inline ::flyteidl::core::ContainerError* ErrorDocument::mutable_error() {
  
  if (error_ == NULL) {
    error_ = new ::flyteidl::core::ContainerError;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.core.ErrorDocument.error)
  return error_;
}
inline void ErrorDocument::set_allocated_error(::flyteidl::core::ContainerError* error) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == NULL) {
    delete error_;
  }
  if (error) {
    ::google::protobuf::Arena* submessage_arena = NULL;
    if (message_arena != submessage_arena) {
      error = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, error, submessage_arena);
    }
    
  } else {
    
  }
  error_ = error;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.core.ErrorDocument.error)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace core
}  // namespace flyteidl

namespace google {
namespace protobuf {

template <> struct is_proto_enum< ::flyteidl::core::ContainerError_Kind> : ::google::protobuf::internal::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::flyteidl::core::ContainerError_Kind>() {
  return ::flyteidl::core::ContainerError_Kind_descriptor();
}

}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_flyteidl_2fcore_2ferrors_2eproto__INCLUDED