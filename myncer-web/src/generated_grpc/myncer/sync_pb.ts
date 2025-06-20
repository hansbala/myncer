// @generated by protoc-gen-es v2.5.2 with parameter "target=ts"
// @generated from file myncer/sync.proto (package myncer, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv2";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { Datasource } from "./datasource_pb";
import { file_myncer_datasource } from "./datasource_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file myncer/sync.proto.
 */
export const file_myncer_sync: GenFile = /*@__PURE__*/
  fileDesc("ChFteW5jZXIvc3luYy5wcm90bxIGbXluY2VyIr8BCgRTeW5jEgoKAmlkGAEgASgJEg8KB3VzZXJfaWQYAiABKAkSLgoKY3JlYXRlZF9hdBgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXASLgoKdXBkYXRlZF9hdBgEIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXASKgoMb25lX3dheV9zeW5jGAUgASgLMhIubXluY2VyLk9uZVdheVN5bmNIAEIOCgxzeW5jX3ZhcmlhbnQidwoKT25lV2F5U3luYxIjCgZzb3VyY2UYASABKAsyEy5teW5jZXIuTXVzaWNTb3VyY2USKAoLZGVzdGluYXRpb24YAiABKAsyEy5teW5jZXIuTXVzaWNTb3VyY2USGgoSb3ZlcndyaXRlX2V4aXN0aW5nGAMgASgIIkoKC011c2ljU291cmNlEiYKCmRhdGFzb3VyY2UYASABKA4yEi5teW5jZXIuRGF0YXNvdXJjZRITCgtwbGF5bGlzdF9pZBgCIAEoCUIzWjFnaXRodWIuY29tL2hhbnNiYWxhL215bmNlci9wcm90by9teW5jZXI7bXluY2VyX3BiYgZwcm90bzM", [file_google_protobuf_timestamp, file_myncer_datasource]);

/**
 * @generated from message myncer.Sync
 */
export type Sync = Message<"myncer.Sync"> & {
  /**
   * google/uuid generated UUID.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * Myncer user id.
   *
   * @generated from field: string user_id = 2;
   */
  userId: string;

  /**
   * Metadata which is fetched from SQL (for it's ACID compliance).
   *
   * @generated from field: google.protobuf.Timestamp created_at = 3;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 4;
   */
  updatedAt?: Timestamp;

  /**
   * Holds the actual sync data.
   *
   * @generated from oneof myncer.Sync.sync_variant
   */
  syncVariant: {
    /**
     * @generated from field: myncer.OneWaySync one_way_sync = 5;
     */
    value: OneWaySync;
    case: "oneWaySync";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message myncer.Sync.
 * Use `create(SyncSchema)` to create a new message.
 */
export const SyncSchema: GenMessage<Sync> = /*@__PURE__*/
  messageDesc(file_myncer_sync, 0);

/**
 * Representative of source -> destination.
 *
 * @generated from message myncer.OneWaySync
 */
export type OneWaySync = Message<"myncer.OneWaySync"> & {
  /**
   * @generated from field: myncer.MusicSource source = 1;
   */
  source?: MusicSource;

  /**
   * @generated from field: myncer.MusicSource destination = 2;
   */
  destination?: MusicSource;

  /**
   * When true, it overwrites the destination songs.
   * If a song exists in source but not in destination, the song will be lost from destination.
   *
   * next: 4
   *
   * @generated from field: bool overwrite_existing = 3;
   */
  overwriteExisting: boolean;
};

/**
 * Describes the message myncer.OneWaySync.
 * Use `create(OneWaySyncSchema)` to create a new message.
 */
export const OneWaySyncSchema: GenMessage<OneWaySync> = /*@__PURE__*/
  messageDesc(file_myncer_sync, 1);

/**
 * @generated from message myncer.MusicSource
 */
export type MusicSource = Message<"myncer.MusicSource"> & {
  /**
   * @generated from field: myncer.Datasource datasource = 1;
   */
  datasource: Datasource;

  /**
   * Unique, stable playlist identifier for the datasource.
   *
   * next: 3
   *
   * @generated from field: string playlist_id = 2;
   */
  playlistId: string;
};

/**
 * Describes the message myncer.MusicSource.
 * Use `create(MusicSourceSchema)` to create a new message.
 */
export const MusicSourceSchema: GenMessage<MusicSource> = /*@__PURE__*/
  messageDesc(file_myncer_sync, 2);

