/* tslint:disable */
/* eslint-disable */
/**
 * Myncer API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { SyncSyncData } from './SyncSyncData';
import {
    SyncSyncDataFromJSON,
    SyncSyncDataFromJSONTyped,
    SyncSyncDataToJSON,
    SyncSyncDataToJSONTyped,
} from './SyncSyncData';
import type { SyncVariant } from './SyncVariant';
import {
    SyncVariantFromJSON,
    SyncVariantFromJSONTyped,
    SyncVariantToJSON,
    SyncVariantToJSONTyped,
} from './SyncVariant';

/**
 * 
 * @export
 * @interface Sync
 */
export interface Sync {
    /**
     * Unique id of the sync.
     * @type {string}
     * @memberof Sync
     */
    id: string;
    /**
     * The timestamp this sync was created at.
     * @type {Date}
     * @memberof Sync
     */
    createdAt: Date;
    /**
     * The timestamp this sync was updated at.
     * @type {Date}
     * @memberof Sync
     */
    updatedAt: Date;
    /**
     * 
     * @type {SyncVariant}
     * @memberof Sync
     */
    syncVariant: SyncVariant;
    /**
     * 
     * @type {SyncSyncData}
     * @memberof Sync
     */
    syncData: SyncSyncData;
}



/**
 * Check if a given object implements the Sync interface.
 */
export function instanceOfSync(value: object): value is Sync {
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('createdAt' in value) || value['createdAt'] === undefined) return false;
    if (!('updatedAt' in value) || value['updatedAt'] === undefined) return false;
    if (!('syncVariant' in value) || value['syncVariant'] === undefined) return false;
    if (!('syncData' in value) || value['syncData'] === undefined) return false;
    return true;
}

export function SyncFromJSON(json: any): Sync {
    return SyncFromJSONTyped(json, false);
}

export function SyncFromJSONTyped(json: any, ignoreDiscriminator: boolean): Sync {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'],
        'createdAt': (new Date(json['createdAt'])),
        'updatedAt': (new Date(json['updatedAt'])),
        'syncVariant': SyncVariantFromJSON(json['syncVariant']),
        'syncData': SyncSyncDataFromJSON(json['syncData']),
    };
}

export function SyncToJSON(json: any): Sync {
    return SyncToJSONTyped(json, false);
}

export function SyncToJSONTyped(value?: Sync | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'createdAt': ((value['createdAt']).toISOString()),
        'updatedAt': ((value['updatedAt']).toISOString()),
        'syncVariant': SyncVariantToJSON(value['syncVariant']),
        'syncData': SyncSyncDataToJSON(value['syncData']),
    };
}

